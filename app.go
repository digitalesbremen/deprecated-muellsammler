package main

import (
	"bremen_trash/net/http"
	xml2 "bremen_trash/xml"
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
	A       A      `xml:"a"`
	Class   string `xml:"class,attr"`
	OnClick string `xml:"onClick,attr"`
}

type A struct {
	Href string `xml:"href,attr"`
	A    string `xml:",innerxml"`
}

func main() {
	content, err := http.GetContent(root)

	if err != nil {
		log.Fatal(err)
	}

	content = xml2.RepairInvalidHtml(content)
	fmt.Println(content)

	content = strings.ReplaceAll(content, "<br>", "")
	content = strings.ReplaceAll(content, "</br>", "")

	decoder := xml.NewDecoder(strings.NewReader(content))
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
				err = decoder.DecodeElement(&td, &startElement)
				fmt.Println(err)
				fmt.Println("a", td.A)
				fmt.Println("Start: ", startElement.Name.Local, startElement.Attr)
			}
		case xml.EndElement:
			continue
			//fmt.Println("End: ", t.Name)
		}
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
