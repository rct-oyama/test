package main

import (
        "fmt"
        "math/rand"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	v := rand.Intn(6) + 1
	s := fmt.Sprintf("サイコロの目は、%d", v)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: s,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
