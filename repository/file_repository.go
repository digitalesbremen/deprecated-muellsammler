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

func write() {
	data := Addresses{
		Addresses: []Address{
			{
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
			}, {
				Street:      "Teststr.",
				HouseNumber: "17",
				CollectionDates: []GarbageCollectionDate{
					{
						Date:  JSONTime{time.Date(2020, 05, 23, 0, 0, 0, 0, time.UTC)},
						Types: []string{"YellowBag", "PaperWaste"},
					},
				},
			},
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("../test.json", file, 0644)

	readFile, _ := ioutil.ReadFile("../test.json")

	var addresses Addresses
	_ = json.Unmarshal(readFile, &addresses)
	fmt.Println(addresses)
}
