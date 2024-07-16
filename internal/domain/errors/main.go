package domainerrors

import "errors"

var (
	RequiredProperty = errors.New("missing required property")
	EntityNotFound   = errors.New("entity not found")
)
