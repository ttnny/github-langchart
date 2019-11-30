package main

import (
	"html/template"
	"net/http"
)

type LangChart struct {
	Langs     []string
	Stats     []int
	ChartType string
}

// Route: (Demo) Handle Index
func demoIndexHandleFunc(w http.ResponseWriter, _ *http.Request) {
}

// Route: (Demo) Handle Display GitHub LangStats
func demoLangStatsHandleFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/langstats.gohtml", "templates/langchart.gohtml"))

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		langStats := getLangStats(username)

		// Convert data from map to struct
		langs := []string{}
		stats := []int{}
		for l, s := range langStats {
			langs = append(langs, l)
			stats = append(stats, s)
		}

		// Prepare dataset for the chart
		var langChart LangChart
		langChart.Langs = langs
		langChart.Stats = stats
		langChart.ChartType = "bar"

		tmpl.Execute(w, langChart)
	}
}

// Route: (Demo) Handle Display GitHub CtbnStats
func demoCtbnStatsHandleFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/ctbnstats.gohtml"))

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		svg := getCtbnStats(username)

		tmpl.Execute(w, template.HTML(svg))
	}
}
