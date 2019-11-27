package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v28/github"
	"github.com/gorilla/mux"
	"github.com/wcharczuk/go-chart"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	// Route: index
	//r.HandleFunc("/", indexHandleFunc).Methods("GET")

	// Route: chart
	//r.HandleFunc("/app/github-lcg/api/chart/{username}", chartHandleFunc).Methods("GET")
	rank := getLangRanking("ttnny")
	for key, value := range rank {
		fmt.Println(key, value)
	}

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
	_, err := fmt.Fprintf(w, "Hi there! Welcome to GitHub Languages Chart Generator.")
	if err != nil {
		fmt.Printf("Errors: %v\n", err)
	}
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

// Create pie chart
func createPieChart() {
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}

// Calculate GitHub language ranking
func getLangRanking(username string) map[string]int {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "a13d09c261f1e95c3de70e040d5f35c0a405226c"},
	)

	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

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
	langs := make(map[string]int)
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

	return langs
}

// Get/set the default port
func port() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}