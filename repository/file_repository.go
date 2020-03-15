package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Marshaler interface {
	MarshalJSON() ([]byte, error)
	UnmarsalJSON(b []byte) error
}

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
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

func Save(data Addresses, fullQualifiedFileName string) {
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(fullQualifiedFileName, file, 0644)
}
