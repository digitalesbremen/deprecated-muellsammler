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

	for _, element := range firstLetters {
		fmt.Println(element)
	}
}
