package xlambda

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// NewProxyRequest is a helper method to build a events.APIGatewayProxyRequest object.
// You can use this method in tests to mock incoming request payloads to a Lambda function.
func NewProxyRequest(method string, queryParameters map[string]string, body interface{}) (*events.APIGatewayProxyRequest, error) {
	request := &events.APIGatewayProxyRequest{
		HTTPMethod:            method,
		QueryStringParameters: queryParameters,
	}

	if body == nil {
		return request, nil
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request.Body = string(data)

	return request, nil
}
