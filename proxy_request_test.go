package xlambda_test

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"

	"github.com/strongishllama/xlambda"
)

func TestNewProxyRequest(t *testing.T) {
	request, err := xlambda.NewProxyRequest(http.MethodGet, map[string]string{
		"test-key": "test-value",
	}, struct {
		Key string `json:"key"`
	}{"value"})

	require.NoError(t, err)
	require.Equal(t, &events.APIGatewayProxyRequest{
		HTTPMethod: http.MethodGet,
		QueryStringParameters: map[string]string{
			"test-key": "test-value",
		},
		Body: `{"key":"value"}`,
	}, request)
}
