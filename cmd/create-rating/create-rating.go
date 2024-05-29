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

	var reviewDto database.ReviewDto
	if err := request.JSON(req, &reviewDto); err != nil {
		return response.Text(http.StatusBadRequest, err.Error())
	}
	if !reviewDto.IsRatingValid() {
		return response.Text(http.StatusBadRequest, "rating must be 1-10")
	}

	review := reviewDto.ToReview()
	if err := db.Create(&review).Error; err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return response.JSON(http.StatusOK, review)
}

func main() {
	db = database.Must(database.Connect())
	lambda.Start(handler)
}
