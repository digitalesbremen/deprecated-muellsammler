package main

import (
	"bremen_trash/html/bremen/stadtreinigung"
	"bremen_trash/html/repair"
	"bremen_trash/net/http"
	"fmt"
	"log"
	"strings"
)

var (
	bremerStadtreinigungRootUrl  = "http://213.168.213.236/bremereb/bify/"
	bremerStadtreinigungIndexUrl = bremerStadtreinigungRootUrl + "index.jsp"
)

func main() {
	content, err := http.GetContent(bremerStadtreinigungIndexUrl)

	if err != nil {
		log.Fatal(err)
	}

	content = repair.RepairInvalidHtml(content)

	// Hack: Don't now why but parsing </br> does not work.
	content = strings.ReplaceAll(content, "<br>", "")
	content = strings.ReplaceAll(content, "</br>", "")

	firstLetters := stadtreinigung.ParseIndexPage(content, bremerStadtreinigungRootUrl)

	streets := make([]stadtreinigung.Street, 0)

	for _, firstLetter := range firstLetters {
		content, err := http.GetContent(firstLetter.Url)

		if err != nil {
			log.Fatal(err)
		}

		firstLetterStreets, err := stadtreinigung.ParseStreetPage(content, firstLetter)

		if err != nil {
			fmt.Printf(`Error while parsing streets of %s. Error is %s. Url will be ignored.`, firstLetter.Url, err)
		}

		for _, element := range firstLetterStreets {
			streets = append(streets, element)
		}
	}

	for _, street := range streets {
		fmt.Println(`Street name`, street.Name)
		fmt.Println(`Street url`, street.Url)
	}
}
