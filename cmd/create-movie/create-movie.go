package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"gorm.io/gorm"
	"net/http"
	"serverless-rotten-bananas/pkg/database"
	"serverless-rotten-bananas/pkg/request"
	"serverless-rotten-bananas/pkg/response"
)

var db *gorm.DB

func handler(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	if req.RequestContext.HTTP.Method != "POST" {
		return response.Status(http.StatusMethodNotAllowed)
	}

	contentType, ok := req.Headers["content-type"]
	if !ok || contentType != "application/json" {
		return response.Text(http.StatusUnsupportedMediaType, "Content-Type must be application/json")
	}

	var movieDto database.MovieDto
	if err := request.JSON(req, &movieDto); err != nil {
		return response.Text(http.StatusBadRequest, err.Error())
	}

	movie := movieDto.ToMovie()
	if err := db.Create(&movie).Error; err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return response.JSON(http.StatusOK, movie)
}

func main() {
	db = database.Must(database.Connect())
	lambda.Start(handler)
}
