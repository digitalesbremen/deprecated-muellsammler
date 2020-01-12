package stadtreinigung

import (
	"bremen_trash/html/bremen/stadtreinigung/parser"
	"encoding/xml"
)

func ParseIndexPage(content string, bremerStadtreinigungRootUrl string) []parser.Dto {
	matcher := func(startElement xml.StartElement) bool { return matchesFirstLetterOfStreet(startElement) }
	return parser.ParseHtml(content, matcher, bremerStadtreinigungRootUrl)
}

func matchesFirstLetterOfStreet(startElement xml.StartElement) bool {
	if startElement.Name.Local == `td` {
		for _, attribute := range startElement.Attr {
			if attribute.Name.Local == `class` && attribute.Value == `BAKChr` {
				return true
			}
		}
	}

	return false
}
