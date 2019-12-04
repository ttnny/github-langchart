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
	r.HandleFunc("/", demoIndexHandleFunc).Methods("GET")

	// Route: (Demo) Display GitHub LangStats
	r.HandleFunc("/github-lcs/langstats", demoLangStatsHandleFunc)

	// Route: (Demo) Display GitHub CtbnStats
	r.HandleFunc("/github-lcs/ctbnstats", demoCtbnStatsHandleFunc)

	// Serve static files
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("static"))))

	// Let's start
	err := http.ListenAndServe(port(), r)
	if err != nil {
		fmt.Printf("Errors: %v\n", err)
	}
}

// Get/set the default port
func port() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}