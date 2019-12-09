package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"os"
)

var errLog = log.New(os.Stderr, "ERROR ", log.Llongfile)

func main() {
	lambda.Start(ctbnStatsHandleFunc)
}

// Handle API: Github Contribution Statistics
func ctbnStatsHandleFunc(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the path parameter {username} from the request
	username := r.PathParameters["username"]

	// Get the contribution graph
	graph := getCtbnStats(username)

	// Prepare SVG format
	svg := graph[:4] + ` version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" style="font-size: 10px; fill: #767676; font-family: sans-serif, Arial, Helvetica;"` + graph[4:]

	// Return a response with a 200 OK status
	// and the SVG (string) in the body.
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       svg,
		Headers: map[string]string{
			"Content-Type": "image/svg+xml",
		},
	}, nil
}

// Get GitHub contribution graph
func getCtbnStats(username string) string {
	url := "https://github.com/users/" + username + "/contributions"

	// Request the HTML page
	res, err := http.Get(url)
	if err != nil {
		_, _ = clientError(http.StatusNotFound)
		return ""
	}

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Get the SVG
	s := doc.Find("svg.js-calendar-graph-svg")
	svg, err := goquery.OuterHtml(s)
	if err != nil {
		log.Fatal(err)
	}

	return svg
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
