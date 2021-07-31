package xlambda

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gofor-little/xerror"
)

// UnmarshalAndValidate unmarshals the request's body into v and then calls Validate on it.
// v must be a pointer to an object and cannot be nil.
func UnmarshalAndValidate(request *events.APIGatewayProxyRequest, v Validatable) error {
	if err := json.Unmarshal([]byte(request.Body), v); err != nil {
		return xerror.Wrap("failed to unmarshal request body into Validatable", err)
	}

	if err := v.Validate(); err != nil {
		return xerror.Wrap("failed to validate Validatable", err)
	}

	return nil
}
