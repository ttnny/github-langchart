package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	// Route: index
	r.HandleFunc("/", indexHandleFunc).Methods("GET")

	// Route: chart
	r.HandleFunc("/api/github-lcg/chart/{username}", chartHandleFunc).Methods("GET")

	// Route: contributions
	r.HandleFunc("/api/github-lcg/contributions/{username}", contributionsHandleFunc).Methods("GET")

	// Let's start
	http.ListenAndServe(port(), r)
}

func indexHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there! Welcome to GitHub Languages Chart Generator.")
}

func chartHandleFunc(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	username := vars["username"]

	// Construct a new GitHub client
	gc := github.NewClient(nil)
	result, _, _ := gc.Repositories.ListLanguages(ctx, username, "go")

	fmt.Fprintf(w, "Languages used in the repo 'go' from %s", username)
	fmt.Println(result)
}

func contributionsHandleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", username)
}

func port() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}