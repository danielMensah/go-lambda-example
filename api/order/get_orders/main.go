package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/danielMensah/go-lambda-example/internal/utility/aws"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(req aws.Request) (aws.Response, error) {
	return aws.Handle(req, GetOrders), nil
}

func main() {
	lambda.Start(Handler)
}
