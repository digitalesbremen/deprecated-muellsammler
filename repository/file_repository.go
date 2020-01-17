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
	UnmarshalJSON(b []byte) error
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

type Address struct {
	Street          string
	HouseNumber     string
	CollectionDates []GarbageCollectionDate
}

type GarbageCollectionDate struct {
	Date  JSONTime
	Types []string
}

func write() {
	data := Address{
		Street:      "Teststr.",
		HouseNumber: "15a",
		CollectionDates: []GarbageCollectionDate{
			{
				Date:  JSONTime{time.Date(2018, 07, 05, 0, 0, 0, 0, time.UTC)},
				Types: []string{"YellowBag", "PaperWaste"},
			},
			{
				Date:  JSONTime{time.Date(2020, 01, 11, 0, 0, 0, 0, time.UTC)},
				Types: []string{"ChristmasTree"},
			},
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("../test.json", file, 0644)

	readFile, _ := ioutil.ReadFile("../test.json")

	var address Address
	_ = json.Unmarshal(readFile, &address)
	fmt.Println(address)
}
