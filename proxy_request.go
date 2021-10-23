package xlambda

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/mitchellh/mapstructure"
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

// ParseAndValidate unmarshals the query string parameters into validatable and then calls
// Validate on it. validatable must be a pointer to an object and cannot be nil.
func ParseAndValidate(request *events.APIGatewayProxyRequest, validatable Validatable) error {
	if err := mapstructure.Decode(request.QueryStringParameters, &validatable); err != nil {
		return fmt.Errorf("failed to decode query string parameters: %w", err)
	}

	if err := validatable.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}

// UnmarshalAndValidate unmarshals the request's body into validatable and then calls Validate
// on it. validatable must be a pointer to an object and cannot be nil.
func UnmarshalAndValidate(request *events.APIGatewayProxyRequest, validatable Validatable) error {
	if err := json.Unmarshal([]byte(request.Body), validatable); err != nil {
		return fmt.Errorf("failed to unmarshal request body into Validatable: %w", err)
	}

	if err := validatable.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}
