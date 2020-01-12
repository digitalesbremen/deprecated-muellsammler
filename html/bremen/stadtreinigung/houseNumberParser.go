package stadtreinigung

import (
	"bremen_trash/html/bremen/stadtreinigung/parser"
	"encoding/xml"
	"strings"
)

func ParseHouseNumber(content string, bremerStadtreinigungRootUrl string) []parser.Dto {
	matcher := func(startElement xml.StartElement) bool { return matchesHouseNumber(startElement) }
	return parser.ParseHtml(content, matcher, bremerStadtreinigungRootUrl)
}

func matchesHouseNumber(startElement xml.StartElement) bool {
	if startElement.Name.Local == `td` {
		for _, attribute := range startElement.Attr {
			if attribute.Name.Local == `class` && strings.Contains(attribute.Value, `BAKStr`) {
				return true
			}
		}
	}

	return false
}
