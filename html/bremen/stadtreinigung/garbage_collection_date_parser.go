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
	Type string
}

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

					dates = append(dates, GarageCollection{submatch[1] + `.` + actualYear, submatch[2]})
				}
			}
		}
	}

	return dates
}
