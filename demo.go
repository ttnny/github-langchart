package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Route: (Demo) Handle Index
func demoIndexHandleFunc(w http.ResponseWriter, _ *http.Request) {

}

// Route: (Demo) Handle Display GitHub LangStats
func demoLangStatsHandleFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/langstats.gohtml"))

	if r.Method == http.MethodGet {
		tmpl.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")

		//svg := getLangStats(username)

		//tmpl.Execute(w, template.HTML(svg))
	}

	//rank := getLangStats(username)

	//for key, value := range rank {
	//	fmt.Println(key, value)
	//}

	//toJSON(rank)
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
