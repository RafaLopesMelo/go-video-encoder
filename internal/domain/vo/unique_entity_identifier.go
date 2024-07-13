package vo

import "github.com/google/uuid"

type UniqueEntityID struct {
	value string
}

func (id *UniqueEntityID) Value() string {
	return id.value
}

func NewIDFromValue(id string) *UniqueEntityID {
	return &UniqueEntityID{
		value: id,
	}
}

func NewID() *UniqueEntityID {
	value, _ := uuid.NewV7()

	return &UniqueEntityID{
		value: value.String(),
	}
}
