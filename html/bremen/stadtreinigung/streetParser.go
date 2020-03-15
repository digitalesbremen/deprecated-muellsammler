package stadtreinigung

import (
	"encoding/xml"
	"regexp"
	"strings"

	"muellsammler/html/bremen/stadtreinigung/parser"
)

func ParseStreetPage(content string, firstLetter parser.Dto, bremerStadtreinigungRootUrl string) []parser.Dto {
	streets := make([]parser.Dto, 0)

	if strings.Contains(content, `<!-- BEGIN: Keine Strassen gefunden:-->`) {
		//fmt.Println("Page does not contains streets")
		// do nothing
	} else if strings.Contains(content, `bitte w&auml;hlen Sie die Hausnummer:`) {
		regex := regexp.MustCompile(`<!-- BEGIN: Strassen gefunden:-->\s*<h[0-9]>([` + firstLetter.Value + `][a-zA-Z-]*)<\/h[0-9]>`)
		submatch := regex.FindStringSubmatch(content)
		streetName := submatch[1]
		streetUrl := firstLetter.Url
		street := parser.Dto{streetName, streetUrl}
		streets = append(streets, street)
	} else {
		matcher := func(startElement xml.StartElement) bool { return matchesStreet(startElement) }
		return parser.ParseHtml(content, matcher, bremerStadtreinigungRootUrl)
	}

	return streets
}

func matchesStreet(startElement xml.StartElement) bool {
	if startElement.Name.Local == `td` {
		for _, attribute := range startElement.Attr {
			if attribute.Name.Local == `class` && strings.Contains(attribute.Value, `BAKStr`) {
				return true
			}
		}
	}

	return false
}
