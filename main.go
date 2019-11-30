package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	// Route: (API) Github LangStats (language rankings)
	r.HandleFunc("/github-lcs/api/langstats/{username}", langStatsHandleFunc).Methods("GET")

	// Route: (API) GitHub CtbnStats (contribution graph)
	r.HandleFunc("/github-lcs/api/ctbnstats/{username}", ctbnStatsHandleFunc).Methods("GET")

	// Route: (Demo) Index
	r.HandleFunc("/github-lcs", demoIndexHandleFunc).Methods("GET")

	// Route: (Demo) Display GitHub LangStats
	r.HandleFunc("/github-lcs/displayLangStats", demoLangStatsHandleFunc)

	// Route: (Demo) Display GitHub CtbnStats
	r.HandleFunc("/github-lcs/displayCtbnStats", demoCtbnStatsHandleFunc)

	// Let's start
	err := http.ListenAndServe(port(), r)
	if err != nil {
		fmt.Printf("Errors: %v\n", err)
	}
}

// Handle Route: (API) Github LangStats (language rankings)
func langStatsHandleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	rank := getLangRanking(username)
	for key, value := range rank {
		fmt.Fprintln(w, key, value)
	}
}

// Handle Route: (API) GitHub CtbnStats (contribution graph)
func ctbnStatsHandleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", username)
}

// Get/set the default port
func port() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}