package main

import (
	"fmt"
	"log"

	"bremen_trash/client"
	"bremen_trash/html/bremen/stadtreinigung"
	"bremen_trash/html/bremen/stadtreinigung/parser"
	"bremen_trash/progressbar"
)

var (
	bremerStadtreinigungRootUrl  = "http://213.168.213.236/bremereb/bify/"
	bremerStadtreinigungIndexUrl = bremerStadtreinigungRootUrl + "index.jsp"
	c                            = client.NewClient()
)

func main() {
	// Load first letters
	fmt.Println("Loading street first letters")
	bar := progressbar.NewProgressBar(1)
	firstLetters := loadFirstLetterOfStreets()
	bar.Finish("%d street first letters loaded", len(firstLetters))

	// Load streets
	fmt.Println("Loading streets")
	bar = progressbar.NewProgressBar(3700)
	streets := loadStreets(firstLetters, bar)
	bar.Finish("%d streets loaded", len(streets))

	type GarbageCollectionUrl struct {
		Street               string
		HouseNumber          string
		GarbageCollectionUrl string
	}

	fmt.Println("Loading house numbers")
	garbageCollectionUrls := make([]GarbageCollectionUrl, 0)
	bar = progressbar.NewProgressBar(len(streets))
	for _, street := range streets {
		content, err := c.GetContent(street.Url)

		if err != nil {
			log.Fatal(err)
		}

		houseNumbers := stadtreinigung.ParseHouseNumber(content, bremerStadtreinigungRootUrl)

		for _, houseNumber := range houseNumbers {
			garbageCollectionUrls = append(garbageCollectionUrls, GarbageCollectionUrl{street.Value, houseNumber.Value, houseNumber.Url})
		}

		bar.Add(1)
		//loadDates(houseNumbers, content)
	}

	bar.Finish("%d house numbers loaded", len(garbageCollectionUrls))

	fmt.Println("Loading garbage collection dates")
	bar = progressbar.NewProgressBar(len(garbageCollectionUrls))
	dates := 0
	for _, garbageCollectionUrl := range garbageCollectionUrls {
		gcd := loadDates(garbageCollectionUrl.GarbageCollectionUrl)
		bar.Add(1)
		dates = dates + len(gcd)
	}
	bar.Finish("%d garbage collection dates loaded", dates)
}

func loadDates(garbageCollectionUrl string) []stadtreinigung.GarageCollection {
	garbageContent, err := c.GetContent(garbageCollectionUrl)

	if err != nil {
		log.Fatal(err)
	}

	if garbageContent == "" {
		log.Fatal(`Dates content is empty. Url: `, garbageCollectionUrl)
	}

	return stadtreinigung.ParseGarbageCollectionDates(garbageContent)
}

func loadStreets(firstLetters []parser.Dto, bar *progressbar.ProgressBar) []parser.Dto {
	streets := make([]parser.Dto, 0)

	for _, firstLetter := range firstLetters {
		//fmt.Println(`Found first letter of street`, firstLetter.FirstLetter, firstLetter.Url)

		content, err := c.GetContent(firstLetter.Url)

		if err != nil {
			log.Fatal(err)
		}

		firstLetterStreets := stadtreinigung.ParseStreetPage(content, firstLetter, bremerStadtreinigungRootUrl)

		for _, element := range firstLetterStreets {
			streets = append(streets, element)
		}

		bar.Add(len(firstLetterStreets))
	}
	return streets
}

func loadFirstLetterOfStreets() []parser.Dto {
	content, err := c.GetContent(bremerStadtreinigungIndexUrl)

	if err != nil {
		log.Fatal(err)
	}

	firstLetters := stadtreinigung.ParseIndexPage(content, bremerStadtreinigungRootUrl)
	return firstLetters
}
