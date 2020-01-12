package stadtreinigung

import (
	"log"
	"regexp"
)

type GarageCollection struct {
	Date string
	Type string
}

var (
	regex = regexp.MustCompile(`<nobr>.*([0-9]{2}.[0-9]{2}.)&nbsp;(.*)<nobr>`)
)

func ParseGarbageCollectionDates(content string) []GarageCollection {
	dates := make([]GarageCollection, 0)

	matches := regex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) != 3 {
			log.Fatal("Match size does not match", match)
		}

		dates = append(dates, GarageCollection{match[1], match[2]})
	}

	return dates
}
