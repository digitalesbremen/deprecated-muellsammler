package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	return c.getContent(url, 10, 0)
}

func (c *Client) getContent(url string, maxRetries int, actualRetry int) (content string, err error) {
	client := http.Client{
		Timeout: c.Timeout,
	}

	resp, err := client.Get(url)

	if err != nil {
		if actualRetry < maxRetries {
			actualRetry++
			fmt.Println()
			fmt.Printf(`Timeout while loading '%s'. Retry in 10 seconds. Try %d of %d`, url, actualRetry, maxRetries)
			fmt.Println()
			time.Sleep(c.RetryTimeAfterTimeout)
			return c.getContent(url, maxRetries, actualRetry)
		} else {
			return "", err
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		reader := charmap.Windows1252.NewDecoder().Reader(resp.Body)

		body, err := ioutil.ReadAll(reader)

		if err != nil {
			if actualRetry < maxRetries {
				actualRetry++
				fmt.Println()
				fmt.Printf(`Error while reading body. Retry in 10 seconds. Try %d of %d`, actualRetry, maxRetries)
				fmt.Println()
				time.Sleep(c.RetryTimeAfterTimeout)
				return c.getContent(url, maxRetries, actualRetry)
			} else {
				return "", err
			}
		}

		content := string(body)
		content = repair.RepairInvalidHtml(content)

		return content, nil
	} else {
		return "", fmt.Errorf("Try to load `%s`. Response code is '%s'", url, resp.Status)
	}
}
