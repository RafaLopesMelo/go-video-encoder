package domainerrors

import "errors"

var (
	RequiredProperty   = errors.New("Missing required property")
	RepositoryNotFound = errors.New("Repository could not found entity")
)
