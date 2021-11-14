package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.naver.com",
		"tmp",
	}

	c := make(chan requestResult)

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for range urls {
		result := <-c
		fmt.Println(result.url, result.status)
	}
}

func hitUrl(url string, c chan<- requestResult) {
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		c <- requestResult{url: url, status: "FAILED"}
	} else {
		c <- requestResult{url: url, status: "OK"}
	}
}
