package main

import (
	"bremen_trash/net/http"
	"fmt"
	"log"
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
}
