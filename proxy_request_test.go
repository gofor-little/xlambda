package xlambda_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"

	"github.com/gofor-little/xlambda"
)

func TestProxyRequest(t *testing.T) {
	request, err := xlambda.ProxyRequest(http.MethodGet, map[string]string{
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

func TestParseAndValidate(t *testing.T) {
	require.NoError(t, xlambda.ParseAndValidate(&events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{
			"string": "test-string",
		},
	}, &parseAndValidate{}))
}

type parseAndValidate struct {
	String string `mapstructure:"string"`
}

func (v parseAndValidate) Validate() error {
	if v.String == "" {
		return errors.New("string cannot be empty")
	}
	return nil
}

func TestUnmarshalAndValidate(t *testing.T) {
	require.NoError(t, xlambda.UnmarshalAndValidate(&events.APIGatewayProxyRequest{
		Body: `{"string": "test-string", "int": 1, "float": 1.1, "error": null}`,
	}, &unmarshalAndValidate{}))
}

type unmarshalAndValidate struct {
	String string  `json:"string"`
	Int    int     `json:"int"`
	Float  float32 `json:"float"`
	Error  error   `json:"error"`
}

func (u *unmarshalAndValidate) Validate() error {
	if u.String == "" {
		return errors.New("string cannot be empty")
	}
	if u.Int == 0 {
		return errors.New("int cannot be 0")
	}
	if u.Float == 0 {
		return errors.New("float cannot be 0")
	}
	if u.Error != nil {
		return errors.New("error must be nil")
	}
	return nil
}
