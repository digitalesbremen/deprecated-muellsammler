package stadtreinigung

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

type GarageCollection struct {
	Date string
	Type string
}

var (
	regex = regexp.MustCompile(`<nobr>.*([0-9]{2}.[0-9]{2}.)&nbsp;(.*)</nobr>`)
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

	var b B

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
			if b.Value == `` {
				if startElement.Name.Local == `b` {
					_ = decoder.DecodeElement(&b, &startElement)
					if !regexp.MustCompile(`[0-9]{4}`).MatchString(b.Value) {
						b.Value = ``
					}
				}
			} else if startElement.Name.Local == `nobr` {
				var nobr Nobr
				_ = decoder.DecodeElement(&nobr, &startElement)

				fmt.Println(nobr.Value)
			}
		}
	}

	matches := regex.FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		if len(match) != 3 {
			log.Fatal("Match size does not match", match)
		}

		dates = append(dates, GarageCollection{match[1], match[2]})
	}

	return dates
}
