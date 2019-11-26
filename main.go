package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v28/github"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	// Route: index
	r.HandleFunc("/", indexHandleFunc).Methods("GET")

	// Route: chart
	r.HandleFunc("/app/github-lcg/api/chart/{username}", chartHandleFunc).Methods("GET")

	// Route: contributions
	r.HandleFunc("/app/github-lcg/api/contributions/{username}", contributionsHandleFunc).Methods("GET")

	// Let's start
	err := http.ListenAndServe(port(), r)
	if err != nil {
		fmt.Printf("Errors: %v\n", err)
	}
}

// Route: index/home
func indexHandleFunc(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there! Welcome to GitHub Languages Chart Generator.")
	if err != nil {
		fmt.Printf("Errors: %v\n", err)
	}
}

// Route: generates chart
func chartHandleFunc(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	username := vars["username"]

	client := github.NewClient(nil)

	// Get a slice of
	repos, _, err := client.Repositories.List(ctx, username, nil)

	// Address API rate limit and other errors
	if err != nil {
		fmt.Printf("Errors: %v\n", err)
		os.Exit(1)
	}

	// Create a list of repos
	var list []string
	for _, repo := range repos {
		list = append(list, *repo.Name)
	}

	// Get languages info for each repo
	var langs map[string]int
	for _, repo := range list {
		lang, _, err := client.Repositories.ListLanguages(ctx, username, repo)
		if err != nil {
			fmt.Printf("Errors: %v\n", err)
		}

		for k, v := range lang {
			if value, found := langs[k]; found { // if exists, add up value only
				langs[k] = v + value
			} else {
				langs[k] = v
			}
		}
	}

	for key, value := range langs { // Order not specified
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