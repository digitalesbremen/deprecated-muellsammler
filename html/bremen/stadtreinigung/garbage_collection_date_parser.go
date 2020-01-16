package stadtreinigung

import (
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type GarageCollection struct {
	Date string
	Type []Type
}

type Type int

const (
	YellowBag     Type = iota
	ResidualWaste Type = iota
	BioWaste      Type = iota
	PaperWaste    Type = iota
	ChristmasTree Type = iota
)

func ParseGarbageCollectionDates(content string) []GarageCollection {
	dates := make([]GarageCollection, 0)

	decoder := xml.NewDecoder(strings.NewReader(content))
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose

	type B struct {
		Value string `xml:",innerxml"`
	}

	type Nobr struct {
		Value string `xml:",innerxml"`
	}

	actualYear := ``

	for {
		token, tokenErr := decoder.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				break
			}
			fmt.Println(tokenErr)
			break
		}

		switch startElement := token.(type) {
		case xml.StartElement:
			if startElement.Name.Local == `b` {
				var b B
				_ = decoder.DecodeElement(&b, &startElement)
				matchedYear := regexp.MustCompile(`([0-9]{4})`).FindStringSubmatch(b.Value)

				if len(matchedYear) == 2 {
					actualYear = matchedYear[1]
				}
			}

			if actualYear != `` && startElement.Name.Local == `nobr` {
				var nobr Nobr
				_ = decoder.DecodeElement(&nobr, &startElement)

				submatch := regexp.MustCompile(`([0-9]{2}.[0-9]{2})\.&nbsp;(.*)`).FindStringSubmatch(nobr.Value)

				if len(submatch) == 3 {
					fmt.Printf(`%s.%s - %s`, submatch[1], actualYear, submatch[2])
					fmt.Println()

					dates = append(dates, GarageCollection{submatch[1] + `.` + actualYear, mapWasteStrings(submatch[2])})
				}
			}
		}
	}

	return dates
}

func mapWasteStrings(waste string) []Type {
	types := make([]Type, 0)

	if strings.Contains(waste, `Papier`) {
		types = append(types, PaperWaste)
	}

	if strings.Contains(waste, `Gelber Sack`) || strings.Contains(waste, `G.Sack`) {
		types = append(types, YellowBag)
	}

	if strings.Contains(waste, `Tannenbaum`) {
		types = append(types, ChristmasTree)
	}

	if strings.Contains(waste, `Restm`) {
		types = append(types, ResidualWaste)
	}

	if strings.Contains(waste, `Bioabf`) {
		types = append(types, BioWaste)
	}

	return types
}
