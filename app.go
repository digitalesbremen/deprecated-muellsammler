package main

import (
	"bremen_trash/html/bremen/stadtreinigung"
	"bremen_trash/html/bremen/stadtreinigung/parser"
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
	// Load first letters
	fmt.Println("Loading street first letters")
	bar := progressbar.NewOptions(1, progressbar.OptionSetRenderBlankState(true))
	firstLetters := loadFirstLetterOfStreets()
	bar.Finish()
	fmt.Printf("%d street first letters loaded", len(firstLetters))
	fmt.Println()
	fmt.Println()

	// Load streets
	fmt.Println("Loading streets")
	bar = progressbar.NewOptions(3700, progressbar.OptionSetRenderBlankState(true))
	streets := loadStreets(firstLetters, bar)
	bar.Finish()
	fmt.Println()
	fmt.Printf("%d streets loaded", len(streets))
	fmt.Println()
	fmt.Println()

	fmt.Println("Loading house numbers")
	numberOfHouseNumbers := 0
	bar = progressbar.NewOptions(len(streets), progressbar.OptionSetRenderBlankState(true))
	// load house numbers for all streets
	for _, street := range streets {
		content, err := http.GetContent(street.Url)
		content = repair.RepairInvalidHtml(content)

		// Hack: Fix <h3> ends with </h2>
		content = strings.ReplaceAll(content, "<h2>", "<h3>")
		content = strings.ReplaceAll(content, "</h2>", "</h3>")

		if err != nil {
			log.Fatal(err)
		}

		houseNumbers := stadtreinigung.ParseHouseNumber(content, bremerStadtreinigungRootUrl)
		numberOfHouseNumbers = numberOfHouseNumbers + len(houseNumbers)

		bar.Add(1)
		//loadDates(houseNumbers, content)
	}

	bar.Finish()
	fmt.Println()
	fmt.Printf("%d house numbers loaded", numberOfHouseNumbers)
	fmt.Println()
}

func loadDates(houseNumbers []parser.Dto, content string) {
	for _, houseNumber := range houseNumbers {
		garbageContent, err := http.GetContent(houseNumber.Url)
		garbageContent = repair.RepairInvalidHtml(garbageContent)

		if err != nil {
			log.Fatal(err)
		}

		if content == "" {
			log.Fatal(`Dates content is empty. Url: `, houseNumber.Url)
		}

		dates := stadtreinigung.ParseGarbageCollectionDates(garbageContent)

		for _, date := range dates {
			fmt.Println(date.Date, date.Type)
		}
	}
}

func loadStreets(firstLetters []parser.Dto, bar *progressbar.ProgressBar) []parser.Dto {
	streets := make([]parser.Dto, 0)

	for _, firstLetter := range firstLetters {
		//fmt.Println(`Found first letter of street`, firstLetter.FirstLetter, firstLetter.Url)

		content, err := http.GetContent(firstLetter.Url)
		content = repair.RepairInvalidHtml(content)

		if err != nil {
			log.Fatal(err)
		}

		firstLetterStreets := stadtreinigung.ParseStreetPage(content, firstLetter, bremerStadtreinigungRootUrl)

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

func loadFirstLetterOfStreets() []parser.Dto {
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
