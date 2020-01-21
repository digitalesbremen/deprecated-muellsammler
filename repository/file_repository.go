package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const filePath = `../test.json`

type Marshaler interface {
	MarshalJSON() ([]byte, error)
	UnmarsalJSON(b []byte) error
}

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(stamp), nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	parsedDate, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*t = JSONTime{parsedDate}
	return nil
}

type Addresses struct {
	Addresses []Address
}

type Address struct {
	Street          string
	HouseNumber     string
	CollectionDates []GarbageCollectionDate
}

type GarbageCollectionDate struct {
	Date  JSONTime
	Types []string
}

func Save(data Addresses) {
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(filePath, file, 0644)
}

func FindAll() Addresses {
	var addresses Addresses

	readFile, err := ioutil.ReadFile(filePath)

	if err == nil {
		_ = json.Unmarshal(readFile, &addresses)
	}

	return addresses
}
