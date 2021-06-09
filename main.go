package main

import (
        "fmt"
	"strconv"
	"errors"

        //"math/rand"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/osamingo/checkdigit"
)

func getupc(seed string) (string, error) {
	p := checkdigit.NewUPC()
	cd, err := p.Generate(seed)
	if err != nil {
		return "error",errors.New("failed to generate check digit")
	}
	ok := p.Verify(seed + strconv.Itoa(cd))
	return fmt.Sprintf("seed: %s, check digit: %d, verify: %t\n", seed, cd, ok),nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//u, _ := url.Parse("https://xxx.com?a=AAA&b=BBB&c=CCC&d=DDD")
        //query := u.Query()
	//s := fmt.Sprintf("%v\n",request.MultiValueQueryStringParameters)
	
	dgit, _ := getupc(request.MultiValueQueryStringParameters["utid"][0])
	
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: dgit,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
