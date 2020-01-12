package stadtreinigung

import (
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"regexp"
	"strings"
)

type Street struct {
	Name string
	Url  string
}

func ParseStreetPage(content string, firstLetter FirstLetter, bremerStadtreinigungRootUrl string) ([]Street, error) {
	streets := make([]Street, 0)
	var err error = nil

	if strings.Contains(content, `<!-- BEGIN: Keine Strassen gefunden:-->`) {
		err = fmt.Errorf("Page does not contains streets")
	} else if strings.Contains(content, `bitte w&auml;hlen Sie die Hausnummer:`) {
		regex := regexp.MustCompile(`<!-- BEGIN: Strassen gefunden:-->\s*<h[0-9]>([` + firstLetter.FirstLetter + `][a-zA-Z-]*)<\/h[0-9]>`)
		submatch := regex.FindStringSubmatch(content)
		streetName := submatch[1]
		streetUrl := firstLetter.Url
		street := Street{streetName, streetUrl}
		streets = append(streets, street)
	} else {
		decoder := xml.NewDecoder(strings.NewReader(content))
		decoder.Strict = false
		decoder.AutoClose = xml.HTMLAutoClose

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
				if matchesStreet(startElement) {
					var td Td
					err := decoder.DecodeElement(&td, &startElement)

					if err != nil {
						fmt.Printf("Unable to decode tag %s, Tag skipped", startElement.Name.Local)
						continue
					}

					if td.matches() {
						streets = append(streets, td.mapToStreet(bremerStadtreinigungRootUrl))
					}
				}
			}
		}
	}

	return streets, err
}

func (td Td) mapToStreet(bremerStadtreinigungRootUrl string) Street {
	return Street{html.UnescapeString(td.A.Value), bremerStadtreinigungRootUrl + td.A.Href}
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
