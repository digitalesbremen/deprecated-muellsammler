package main

import (
	"bremen_trash/net/http"
	xml2 "bremen_trash/xml"
	"container/list"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"strings"
)

var (
	root = "http://213.168.213.236/bremereb/bify/index.jsp"
)

type Td struct {
	A       *A     `xml:"a"`
	Class   string `xml:"class,attr"`
	OnClick string `xml:"onClick,attr"`
}

type A struct {
	Href  string `xml:"href,attr"`
	Value string `xml:",innerxml"`
}

type FirstLetter struct {
	FirstLetter string
	Url         string
}

func main() {
	content, err := http.GetContent(root)

	if err != nil {
		log.Fatal(err)
	}

	content = xml2.RepairInvalidHtml(content)

	content = strings.ReplaceAll(content, "<br>", "")
	content = strings.ReplaceAll(content, "</br>", "")

	decoder := xml.NewDecoder(strings.NewReader(content))
	decoder.Strict = false
	decoder.AutoClose = xml.HTMLAutoClose

	firstLetters := list.New()

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
				err = decoder.DecodeElement(&td, &startElement)

				if err != nil {
					fmt.Printf("Unable to decode tag %s, Tag skipped", startElement.Name.Local)
					continue
				}

				if td.A != nil && td.A.Href != "" && td.A.Value != "" {
					firstLetters.PushBack(FirstLetter{td.A.Value, td.A.Href})
				}
			}
		case xml.EndElement:
			continue
			//fmt.Println("End: ", t.Name)
		}
	}

	for e := firstLetters.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
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
