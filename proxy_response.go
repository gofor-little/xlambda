package xlambda

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gofor-little/log"
)

// ProxyResponseHTML builds an API gateway proxy response where the body's content type is text/html.
// statusCode should be a valid HTTP status code.
// If err is nil no error will be returned.
// If data is nil nothing will be written to the response body.
func ProxyResponseHTML(statusCode int, err error, data interface{}) (*events.APIGatewayProxyResponse, error) {
	if err != nil {
		log.Error(log.Fields{
			"error":      err,
			"message":    "api request failed",
			"statusCode": statusCode,
		})
	}

	response := &events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "text/html",
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent",
			"Access-Control-Allow-Origin":  accessControlAllowOrigin,
			"Access-Control-Allow-Methods": "OPTIONS,GET,PUT,POST,DELETE,PATCH,HEAD",
		},
		StatusCode: statusCode,
	}

	if data != nil {
		response.Body = fmt.Sprintf("%s", data)
	}

	return response, nil
}

// ProxyResponseJSON builds an API gateway proxy response where the body's content type is application/json.
// statusCode should be a valid HTTP status code.
// If err is nil no error will be returned.
// If data is nil nothing will be written to the response body.
func ProxyResponseJSON(statusCode int, err error, data interface{}) (*events.APIGatewayProxyResponse, error) {
	if err != nil {
		log.Error(log.Fields{
			"error":      err,
			"message":    "api request failed",
			"statusCode": statusCode,
		})
	}

	response := &events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent",
			"Access-Control-Allow-Origin":  accessControlAllowOrigin,
			"Access-Control-Allow-Methods": "OPTIONS,GET,PUT,POST,DELETE,PATCH,HEAD",
		},
		StatusCode: statusCode,
	}

	if data != nil {
		body, marshalErr := json.Marshal(data)
		if marshalErr != nil {
			log.Error(log.Fields{
				"error":      fmt.Errorf("failed to marshal response body and API request failed: %w", marshalErr),
				"statusCode": statusCode,
			})
			return response, nil
		}

		response.Body = string(body)
	}

	return response, nil
}
