package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest() {
	// return "hello", nil
}

func main() {
	lambda.Start(HandleRequest)
}
