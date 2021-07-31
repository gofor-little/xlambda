package xlambda

// Validatable is an interface for an object that can be validated.
type Validatable interface {
	Validate() error
}
