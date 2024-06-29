package errors

type RequiredPropertyError struct {
    property string
}

func (e *RequiredPropertyError) Error() string {
    return "Missing required property: " + e.property
}

func NewRequiredPropertyError(property string) *RequiredPropertyError {
    return &RequiredPropertyError{
        property: property,
    }
}
