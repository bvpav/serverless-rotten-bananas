package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"serverless-rotten-bananas/pkg/database"
	"serverless-rotten-bananas/pkg/response"
	"strings"
)

var db *gorm.DB

func handler(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	if req.RequestContext.HTTP.Method != "GET" {
		return response.Status(http.StatusMethodNotAllowed)
	}

	var searchTerms []string
	if req.QueryStringParameters["search"] != "" {
		searchTerms = strings.Split(req.QueryStringParameters["search"], ",")
	}

	q := db
	for _, term := range searchTerms {
		unesc, err := url.QueryUnescape(term)
		if err != nil {
			return events.APIGatewayProxyResponse{}, err
		}
		unesc = "%" + strings.ReplaceAll(unesc, "%", "\\%") + "%"
		q = q.Where("title ilike ?", unesc)
	}

	var movies []database.Movie
	if err := q.Find(&movies).Error; err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return response.JSON(http.StatusOK, movies)
}

func main() {
	db = database.Must(database.Connect())
	lambda.Start(handler)
}
