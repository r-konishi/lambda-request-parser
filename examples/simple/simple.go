package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/r-konishi/lambda-request-parser/parser"
)

type HelloRequest struct {
	ID   parser.StringNumber `json:"id" validate:"required,min=1,max=999"`
	Name string              `json:"name" validate:"required"`
}
type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (Response, error) {
	hr := HelloRequest{}
	if req.QueryStringParameters != nil && len(req.QueryStringParameters) > 0 {
		err := parser.QueryStringParametersToStruct(&req.QueryStringParameters, &hr)
		if err != nil {
			return Response{StatusCode: 400}, err
		}
	}
	if req.Body != "" {
		err := parser.RequestBodyToStruct(req.Body, &hr)
		if err != nil {
			return Response{StatusCode: 400}, err
		}
	}
	// if req.HTTPMethod == "GET" {
	// 	err := parser.QueryStringParametersToStruct(&req.QueryStringParameters, &hr)
	// 	if err != nil {
	// 		return Response{StatusCode: 400}, err
	// 	}
	// } else {
	// 	err := parser.RequestBodyToStruct(req.Body, &hr)
	// 	if err != nil {
	// 		return Response{StatusCode: 400}, err
	// 	}
	// }

	errors := parser.GetValidationErrors(&hr)
	if errors != nil {
		return Response{StatusCode: 400}, errors
	}

	fmt.Printf("%#v", hr)

	buf := bytes.Buffer{}

	body, err := json.Marshal(map[string]interface{}{
		"message": fmt.Sprintf("Hello, %s!!", hr.Name),
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode: 200,
		Body:       buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
