package domainerrors

import "errors"

type DomainError error

var (
	RequiredProperty    DomainError = errors.New("missing required property")
	EntityNotFound      DomainError = errors.New("entity not found")
	InvalidResourceKind DomainError = errors.New("invalid resource kind")
)
