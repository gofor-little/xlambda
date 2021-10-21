package xlambda

import "errors"

var (
	accessControlAllowOrigin string
)

func Initialize(_accessControlAllowOrigin string) error {
	if len(_accessControlAllowOrigin) == 0 {
		return errors.New("access control allow origin cannot be empty")
	}

	accessControlAllowOrigin = _accessControlAllowOrigin

	return nil
}
