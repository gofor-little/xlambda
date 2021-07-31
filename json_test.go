package xlambda_test

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gofor-little/xerror"
	"github.com/stretchr/testify/require"
	"github.com/strongishllama/xlambda"
)

func TestUnmarshalAndValidate(t *testing.T) {
	require.NoError(t, xlambda.UnmarshalAndValidate(&events.APIGatewayProxyRequest{
		Body: `{"aString": "test-string", "anInt": 1, "aFloat": 1.1, "anError": null}`,
	}, &unmarshalAndValidateTest{}))
}

type unmarshalAndValidateTest struct {
	AString string  `json:"aString"`
	AnInt   int     `json:"anInt"`
	AFloat  float32 `json:"aFloat"`
	AnError error   `json:"anError"`
}

func (v *unmarshalAndValidateTest) Validate() error {
	if v.AString == "" {
		return xerror.New("AString cannot be empty")
	}

	if v.AnInt == 0 {
		return xerror.New("AnInt cannot be 0")
	}

	if v.AFloat == 0 {
		return xerror.New("AFloat cannot be 0")
	}

	if v.AnError != nil {
		return xerror.New("AnError must be nil")
	}

	return nil
}
