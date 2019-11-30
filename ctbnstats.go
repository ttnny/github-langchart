package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func getContributionGraph(username string) string {
	url := "https://github.com/users/" + username + "/contributions"

	// Request the HTML page
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Get the SVG
	s := doc.Find("svg.js-calendar-graph-svg")
	svg, err := goquery.OuterHtml(s)
	if err != nil {
		log.Fatal(err)
	}

	return svg
}
