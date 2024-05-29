package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"gorm.io/gorm"
	"serverless-rotten-bananas/pkg/database"
)

var db *gorm.DB

func handler() {

}

func main() {
	db = database.Must(database.Connect())
	lambda.Start(handler)
}
