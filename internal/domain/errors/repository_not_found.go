package errors

type EntityNotFoundError struct {}

func (e *EntityNotFoundError) Error() string {
    return "Entity does not found."
}

func NewEntityNotFoundError() *EntityNotFoundError {
    return &EntityNotFoundError{}
}
