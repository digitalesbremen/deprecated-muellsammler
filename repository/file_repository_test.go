package repository

import (
	"testing"
	"time"
)

func TestSave(t *testing.T) {
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

	Save(data)
}

func TestRead(t *testing.T) {
	// TODO implement me!
}
