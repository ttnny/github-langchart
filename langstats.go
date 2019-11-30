package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
	"os"
)

// Get GitHub language rankings
func getLangRanking(username string) map[string]int {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "11a32038cfe50fa74f251f69b89a372eaf9074d4"},
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
	langStats := make(map[string]int)
	for _, repo := range list {
		lang, _, err := client.Repositories.ListLanguages(ctx, username, repo)
		if err != nil {
			fmt.Printf("Errors: %v\n", err)
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
