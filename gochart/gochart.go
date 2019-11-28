// [ARCHIVED]
// This package is to generate GitHub languages chart using go-chart module
package gochart

import (
	"context"
	"fmt"
	"github.com/google/go-github/v28/github"
	"github.com/wcharczuk/go-chart"
	"golang.org/x/oauth2"
	"os"
)

// Create bar chart (for go-chart module)
func createBarChart(username string) {
	var values = getLangRanking(username)

	pie := chart.BarChart{
		Title: "Test Bar Chart",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 200,
			},
		},
		Height:   1024,
		BarWidth: 50,
		Bars:     values,
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	_ = pie.Render(chart.PNG, f)
}

// Calculate GitHub language ranking (for go-chart module)
func getLangRanking(username string) []chart.Value {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
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

	return createValuesForChart(langStats)
}

// Convert GitHub language rankings to a
// specific type that go-chart will read (for go-chart module)
func createValuesForChart(langStats map[string]int) []chart.Value {
	var values []chart.Value

	for l, v := range langStats {
		values = append(values, chart.Value{Label: l, Value: float64(v)})
	}

	return values
}
