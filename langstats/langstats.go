package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

var errLog = log.New(os.Stderr, "ERROR ", log.Llongfile)

func main() {
	lambda.Start(langStatsHandleFunc)
}

// Handle API: Github Language Statistics
func langStatsHandleFunc(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the path parameter {username} from the request
	username := r.PathParameters["username"]

	// Get the language stats
	langStats := getLangStats(username)

	// The APIGatewayProxyResponse.Body needs to be a string,
	// so let's first marshal the data into valid JSON.
	js, err := json.Marshal(langStats)
	if err != nil {
		return serverError(err)
	}

	// Return a response with a 200 OK status
	// and the valid JSON as a string in the body.
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

// Get GitHub language statistics
func getLangStats(username string) map[string]int {
	ctx := context.Background()

	// Create a GitHub authenticated client with your token
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "97d3f1d0962aba6dcb8e79bd35c39ab1b1871309"}, )
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Create a GitHub non-authenticated client
	// client := github.NewClient(nil)

	// Get a list of 100 most recent pushed/updated repos from GitHub account
	listOptions := github.ListOptions{Page: 1, PerPage: 100}
	opt := &github.RepositoryListOptions{ListOptions: listOptions, Sort: "pushed"}
	repos, _, err := client.Repositories.List(ctx, username, opt)

	// Address API rate limit and other errors
	if err != nil {
		_, _ = serverError(err)
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
			_, _ = serverError(err)
		}

		for k, v := range lang {
			if value, found := langStats[k]; found { // if exists, add up value
				langStats[k] = v + value
			} else {
				langStats[k] = v
			}
		}
	}

	return langStats
}

// Helper method for handling server errors
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errLog.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// Helper method for handling client errors
func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}
