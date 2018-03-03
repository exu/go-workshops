// main.go
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Example struct {
	Name        string
	Description string
}

func hello() (Example, error) {
	return Example{"Hoł hoł Lambda!", "This is simple lamda example"}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(hello)
}
