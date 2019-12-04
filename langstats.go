package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v28/github"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

// Handle Route: (API) Github LangStats (language rankings)
func langStatsHandleFunc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	langStats := getLangStats(username)

	// Prepare JSON to respond to the request
	json, err := json.Marshal(langStats)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(json))
}

// Get GitHub language statistics
func getLangStats(username string) map[string]int {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "ce72ca5f59a697660c8a2038df3dd264aea8785d"}, )

	// Create a GitHub authenticated client
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Get a list of repos from GitHub account
	repos, _, err := client.Repositories.List(ctx, username, nil)

	// Address API rate limit and other errors
	if err != nil {
		fmt.Printf("Error(s): %v\n", err)
		os.Exit(1)
	}

	// Convert the list of repos to type string slice
	var list []string
	for _, repo := range repos {
		list = append(list, *repo.Name)
	}

	// Get a sum of languages in all repos
	langStats := make(map[string]int)
	for _, repo := range list {
		lang, _, err := client.Repositories.ListLanguages(ctx, username, repo)
		if err != nil {
			log.Fatal(err)
		}

		for k, v := range lang {
			if value, found := langStats[k]; found { // if exists, add up value only
				langStats[k] = v + value
			} else {
				langStats[k] = v
			}
		}
	}

	return langStats
}
