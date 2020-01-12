package main

import (
	"bremen_trash/html/bremen/stadtreinigung"
	"bremen_trash/html/repair"
	"bremen_trash/net/http"
	"fmt"
	"github.com/schollz/progressbar/v2"
	"log"
	"strings"
)

var (
	bremerStadtreinigungRootUrl  = "http://213.168.213.236/bremereb/bify/"
	bremerStadtreinigungIndexUrl = bremerStadtreinigungRootUrl + "index.jsp"
)

func main() {
	bar := progressbar.New(130000)

	firstLetters := loadFirstLetterOfStreets()
	streets := loadStreets(firstLetters, bar)

	requests := 1 + len(streets)

	for _, street := range streets {
		bar.Add(1)

		content, err := http.GetContent(street.Url)
		content = repair.RepairInvalidHtml(content)

		// Hack: Fix <h3> ends with </h2>
		content = strings.ReplaceAll(content, "<h2>", "<h3>")
		content = strings.ReplaceAll(content, "</h2>", "</h3>")

		if err != nil {
			log.Fatal(err)
		}

		houseNumbers := stadtreinigung.ParseHouseNumber(content, bremerStadtreinigungRootUrl)

		requests = requests + len(houseNumbers)
		bar.Add(len(houseNumbers))
	}

	bar.Finish()
	fmt.Println(`Number of requests`, requests)
}

func loadStreets(firstLetters []stadtreinigung.FirstLetter, bar *progressbar.ProgressBar) []stadtreinigung.Street {
	streets := make([]stadtreinigung.Street, 0)

	for _, firstLetter := range firstLetters {
		//fmt.Println(`Found first letter of street`, firstLetter.FirstLetter, firstLetter.Url)

		content, err := http.GetContent(firstLetter.Url)
		content = repair.RepairInvalidHtml(content)

		if err != nil {
			log.Fatal(err)
		}

		firstLetterStreets, err := stadtreinigung.ParseStreetPage(content, firstLetter, bremerStadtreinigungRootUrl)

		//if err != nil {
		//	fmt.Printf(`Error while parsing streets of %s. Error is '%s'. Url will be ignored.`, firstLetter.Url, err)
		//	fmt.Println()
		//}

		for _, element := range firstLetterStreets {
			streets = append(streets, element)
		}

		bar.Add(len(firstLetterStreets))
	}
	return streets
}

func loadFirstLetterOfStreets() []stadtreinigung.FirstLetter {
	content, err := http.GetContent(bremerStadtreinigungIndexUrl)

	if err != nil {
		log.Fatal(err)
	}

	content = repair.RepairInvalidHtml(content)

	// Hack: Don't now why but parsing </br> does not work.
	content = strings.ReplaceAll(content, "<br>", "")
	content = strings.ReplaceAll(content, "</br>", "")

	firstLetters := stadtreinigung.ParseIndexPage(content, bremerStadtreinigungRootUrl)
	return firstLetters
}
