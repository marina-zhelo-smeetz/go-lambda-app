package main

import (
    "context"
    "github.com/aws/aws-lambda-go/lambda"
    "go-lambda-app/src/handlers"
)

func main() {
    lambda.Start(handlers.HandleAlert)
}