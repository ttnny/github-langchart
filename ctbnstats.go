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

	// Get the graph
	graph := getCtbnStats(username)

	// Prepare SVG format
	svg := graph[:4] + ` version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" style="font-size: 10px; fill: #767676; font-family: sans-serif, Arial, Helvetica;"` + graph[4:]

	// Set response type
	w.Header().Set("Content-Type", "image/svg+xml")

	_, err := fmt.Fprintf(w, svg)
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
		fmt.Printf("Error(s): %v\n", err)
		return ""
	}

	defer res.Body.Close()

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
