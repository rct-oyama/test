package main

import (
        "fmt"
        //"math/rand"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//u, _ := url.Parse("https://xxx.com?a=AAA&b=BBB&c=CCC&d=DDD")
        //query := u.Query()
	s := fmt.Printf("%v\n",request.MultiValueQueryStringParameters)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: s,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
