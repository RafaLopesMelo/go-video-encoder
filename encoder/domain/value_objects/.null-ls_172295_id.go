package value_objects

import uuid "github.com/satori/go.uuid"

type ID struct {
    value string
}

type NewIDFromValue(id string) (*ID) {
    return &ID{
        value: id
    }
}

type NewID() *ID {
    return &ID{
    value = uuid.NewV4().String()
    }
}
