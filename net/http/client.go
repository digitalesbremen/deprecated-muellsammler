package http

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"net/http"
	"time"
)

func GetContent(url string) (content string, err error) {
	//fmt.Printf("Request url `%s`\n", url)

	//time.Sleep(10 * time.Millisecond)
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	resp, err := client.Get(url)

	if err != nil {
		fmt.Println()
		fmt.Printf(`Timeout while loading '%s'. Retry in 10 seconds.`, url)
		fmt.Println()
		time.Sleep(10 * time.Second)

		resp, err = client.Get(url)

		if err != nil {
			return "", err
		}
	}

	defer resp.Body.Close()

	//fmt.Println(resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		reader := charmap.Windows1252.NewDecoder().Reader(resp.Body)

		body, err := ioutil.ReadAll(reader)

		if err != nil {
			return "", err
		}

		return string(body), nil
	} else {
		return "", fmt.Errorf("Try to load `%s`. Response code is '%s'", url, resp.Status)
	}
}
