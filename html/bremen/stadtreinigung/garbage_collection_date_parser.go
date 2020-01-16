package stadtreinigung

import (
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"
)

type GarageCollection struct {
	Date time.Time
	Type []WasteType
}

type WasteType int

const (
	YellowBag     WasteType = iota
	ResidualWaste WasteType = iota
	BioWaste      WasteType = iota
	PaperWaste    WasteType = iota
	ChristmasTree WasteType = iota
)

var (
	yearHtmlTagRegex       = regexp.MustCompile(`([0-9]{4})`)
	wasteEntryHtmlTagRegex = regexp.MustCompile(`([0-9]{2}.[0-9]{2})\.&nbsp;(.*)`)
)

func ParseGarbageCollectionDates(content string) []GarageCollection {
	dates := make([]GarageCollection, 0)

	decoder := xml.NewDecoder(strings.NewReader(content))
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose

	type Tag struct {
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
			if startsWithNewYear(startElement) {
				matchedYear := decodeHtmlTagInnerValue(decoder, startElement, yearHtmlTagRegex)

				if len(matchedYear) == 2 {
					actualYear = matchedYear[1]
				}
			}

			if startsWithNewWasteEntry(actualYear, startElement) {
				matchesWasteEntry := decodeHtmlTagInnerValue(decoder, startElement, wasteEntryHtmlTagRegex)

				if len(matchesWasteEntry) == 3 {
					dates = append(dates, GarageCollection{parseDate(matchesWasteEntry[1] + `.` + actualYear), mapWasteStrings(matchesWasteEntry[2])})
				}
			}
		}
	}

	return dates
}

func parseDate(value string) time.Time {
	parse, _ := time.Parse(`02.01.2006`, value)
	return parse
}

func decodeHtmlTagInnerValue(decoder *xml.Decoder, startElement xml.StartElement, regex *regexp.Regexp) []string {
	type Tag struct {
		Value string `xml:",innerxml"`
	}

	var tag Tag
	_ = decoder.DecodeElement(&tag, &startElement)
	return regex.FindStringSubmatch(tag.Value)
}

func startsWithNewYear(startElement xml.StartElement) bool {
	return startElement.Name.Local == `b`
}

func startsWithNewWasteEntry(actualYear string, startElement xml.StartElement) bool {
	return actualYear != `` && startElement.Name.Local == `nobr`
}

func mapWasteStrings(waste string) []WasteType {
	types := make([]WasteType, 0)

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
