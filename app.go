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
	decoder.Entity = xml.HTMLEntity
	decoder.AutoClose = append(decoder.AutoClose, "img")

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

		switch t := token.(type) {
		case xml.StartElement:
			fmt.Println("Start: ", t.Name)
		case xml.EndElement:
			fmt.Println("End: ", t.Name)
		}
	}
}
