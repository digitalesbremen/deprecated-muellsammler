package stadtreinigung

import (
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"strings"
)

type FirstLetter struct {
	FirstLetter string
	Url         string
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

func ParseIndexPage(indexPage string, bremerStadtreinigungRootUrl string) []FirstLetter {
	firstLetters := make([]FirstLetter, 0, 40)

	decoder := xml.NewDecoder(strings.NewReader(indexPage))
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose

	for {
		token, tokenErr := decoder.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				fmt.Println("EOF")
				break
			}
			fmt.Println(tokenErr)
			break
		}

		switch startElement := token.(type) {
		case xml.StartElement:
			if matchesTd(startElement) {
				var td Td
				err := decoder.DecodeElement(&td, &startElement)

				if err != nil {
					fmt.Printf("Unable to decode tag %s, Tag skipped", startElement.Name.Local)
					continue
				}

				if td.A != nil && td.A.Href != "" && td.A.Value != "" {
					firstLetters = append(firstLetters, FirstLetter{html.UnescapeString(td.A.Value), bremerStadtreinigungRootUrl + td.A.Href})
				}
			}
		case xml.EndElement:
			continue
			//fmt.Println("End: ", t.Name)
		}
	}

	return firstLetters
}

func matchesTd(startElement xml.StartElement) bool {
	if startElement.Name.Local == `td` {
		for _, attribute := range startElement.Attr {
			if attribute.Name.Local == `class` && attribute.Value == `BAKChr` {
				return true
			}
		}
	}

	return false
}