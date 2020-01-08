package main

import (
	"bremen_trash/net/http"
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

	fmt.Println(content)

	decoder := xml.NewDecoder(strings.NewReader(content))
	for {
		token, tokenErr := decoder.Token()
		if tokenErr != nil {
			if tokenErr == io.EOF {
				fmt.Println("EOF")
				break
			}
			// handle error
		}

		switch t := token.(type) {
		case xml.StartElement:
			fmt.Println("Start: ", t.Name)
		case xml.EndElement:
			fmt.Println("End: ", t.Name)
		}
	}
}
