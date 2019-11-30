package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Route: (Demo) Handle Index
func demoIndexHandleFunc(w http.ResponseWriter, _ *http.Request) {

}

// Route: (Demo) Handle Display GitHub LangStats
func demoLangStatsHandleFunc(w http.ResponseWriter, r *http.Request) {
	rank := getLangRanking("ttnny")

	for key, value := range rank {
		fmt.Println(key, value)
	}

	toJSON(rank)
}

// Route: (Demo) Handle Display GitHub CtbnStats
func demoCtbnStatsHandleFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/ctbnstats.gohtml"))

	if r.Method == http.MethodPost {
		tmpl.Execute(w, nil)
		username := r.FormValue("username")
		svg := getContributionGraph(username)
		return
	}

	if r.Method == http.MethodGet {

	}

	// tmpl.Execute(w, struct{ Success bool }{true})
}