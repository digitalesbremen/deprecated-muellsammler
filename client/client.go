package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"

	"bremen_trash/client/repair"
)

type Client struct {
	Timeout               time.Duration
	RetryTimeAfterTimeout time.Duration
}

func NewClient() *Client {
	client := Client{
		Timeout:               1 * time.Second,
		RetryTimeAfterTimeout: 10 * time.Second,
	}
	return &client
}

func (c *Client) GetContent(url string) (content string, err error) {
	//fmt.Printf("Request url `%s`\n", url)

	client := http.Client{
		Timeout: c.Timeout,
	}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println()
		fmt.Printf(`Timeout while loading '%s'. Retry in 10 seconds.`, url)
		fmt.Println()
		time.Sleep(c.RetryTimeAfterTimeout)

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

		content := string(body)
		content = repair.RepairInvalidHtml(content)

		// Hack: Don't now why but parsing </br> does not work.
		content = strings.ReplaceAll(content, "<br>", "")
		content = strings.ReplaceAll(content, "</br>", "")

		// Hack: Fix <h3> ends with </h2>
		content = strings.ReplaceAll(content, "<h2>", "<h3>")
		content = strings.ReplaceAll(content, "</h2>", "</h3>")

		return content, nil
	} else {
		return "", fmt.Errorf("Try to load `%s`. Response code is '%s'", url, resp.Status)
	}
}
