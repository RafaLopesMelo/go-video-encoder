package value_objects

import uuid "github.com/satori/go.uuid"

type UniqueEntityID struct {
	value string
}

func NewIDFromValue(id string) *UniqueEntityID {
	return &UniqueEntityID{
		value: id,
	}
}

func NewID() *UniqueEntityID {
	return &UniqueEntityID{
		value: uuid.NewV4().String(),
	}
}
