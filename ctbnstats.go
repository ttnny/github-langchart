package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Handle Route: (API) GitHub CtbnStats (contribution graph)
func ctbnStatsHandleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	_, err := fmt.Fprintln(w, getCtbnStats(username))
	if err != nil {
		log.Fatal(err)
	}
}

// Get GitHub contribution graph
func getCtbnStats(username string) string {
	url := "https://github.com/users/" + username + "/contributions"

	// Request the HTML page
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Status/code error: %d %s", res.StatusCode, res.Status)
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
