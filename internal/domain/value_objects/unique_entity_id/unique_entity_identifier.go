package unique_entity_id

import uuid "github.com/satori/go.uuid"

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
	return &UniqueEntityID{
		value: uuid.NewV4().String(),
	}
}