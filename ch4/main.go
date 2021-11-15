package main

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
)

var baseUrl = "https://kr.indeed.com/jobs?q=python&start=0"

func main() {
	getPages()
}

func getPages() int {
	res, err := http.Get(baseUrl)
	checkErr(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		checkErr(err)
	}(res.Body)

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination")

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
