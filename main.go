package main

import (
	"go-lambda-app/src/handlers"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Printf("Lambda started")
	lambda.Start(handlers.HandleAlert)
}
