package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	// Route: index
	r.HandleFunc("/", indexHandleFunc).Methods("GET")

	// Route: chart
	//r.HandleFunc("/app/github-lcg/api/chart/{username}", chartHandleFunc).Methods("GET")
	rank := getLangRanking("ttnny")
	for key, value := range rank {
		fmt.Println(key, value)
	}

	toJSON(rank)

	// Route: contributions
	//r.HandleFunc("/app/github-lcg/api/contributions/{username}", contributionsHandleFunc).Methods("GET")

	// Let's start
	err := http.ListenAndServe(port(), r)
	if err != nil {
		fmt.Printf("Errors: %v\n", err)
	}
}

// Route: index/home
func indexHandleFunc(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	tmpl.Execute(w, nil)
}

// Route: generates chart
func chartHandleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	rank := getLangRanking(username)
	for key, value := range rank {
		fmt.Fprintln(w, key, value)
	}
}

// Route: generates graph of GitHub contributions
func contributionsHandleFunc(w http.ResponseWriter, r *http.Request) {
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