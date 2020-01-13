package parser

import (
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"strings"
)

type Dto struct {
	Value string
	Url   string
}

type Td struct {
	A       *A     `xml:"a"`
	Class   string `xml:"class,attr"`
	OnClick string `xml:"onClick,attr"`
}

type A struct {
	Href  string `xml:"href,attr"`
	Value string `xml:",innerxml"`
}

// test startElement matches
type matches func(startElement xml.StartElement) bool

func ParseHtml(content string, fn matches, bremerStadtreinigungRootUrl string) []Dto {
	dtos := make([]Dto, 0)

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
			if fn(startElement) {
				var td Td
				err := decoder.DecodeElement(&td, &startElement)

				if err != nil {
					fmt.Printf("Unable to decode tag %s, Tag skipped", startElement.Name.Local)
					continue
				}

				if td.matches() {
					dtos = append(dtos, td.mapToDto(bremerStadtreinigungRootUrl))
				}
			}
		}
	}

	return dtos
}

func (td Td) mapToDto(bremerStadtreinigungRootUrl string) Dto {
	url := bremerStadtreinigungRootUrl + td.A.Href
	url = strings.ReplaceAll(url, ` `, `%20`)
	url = strings.ReplaceAll(url, `ß`, `%DF`)
	return Dto{html.UnescapeString(td.A.Value), url}
}

func (td Td) matches() bool {
	return td.A != nil && td.A.Href != "" && td.A.Value != ""
}
