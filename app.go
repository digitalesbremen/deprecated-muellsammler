package main

import (
	"bremen_trash/net/http/get"
	"fmt"
	"log"
)

var (
	root = "http://213.168.213.236/bremereb/bify/index.jsp"
)

func main() {
	fmt.Println("Hello golang!")

	content, err := get.GetContent(root)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(content)
}
