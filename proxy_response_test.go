package xlambda_test

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
	"github.com/strongishllama/xlambda"
)

func TestNewProxyResponse(t *testing.T) {
	response, err := xlambda.NewProxyResponse(http.StatusOK, xlambda.ContentTypeApplicationJSON, nil, nil)

	require.NoError(t, err)
	require.Equal(t, &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent",
			"Access-Control-Allow-Methods": "OPTIONS,GET,PUT,POST,DELETE,PATCH,HEAD",
			"Access-Control-Allow-Origin":  "",
			"Content-Type":                 "application/json",
		},
	}, response)
}
