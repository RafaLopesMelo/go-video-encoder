package pg

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type jsonable struct {
	value map[string]any
}

func (s *jsonable) Value() (driver.Value, error) {
	return json.Marshal(s.value)
}

func (s *jsonable) Scan(value any) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, &s.value)
	case string:
		return json.Unmarshal([]byte(v), &s.value)
	default:
		return fmt.Errorf("unexpected type %T", v)
	}
}

func (s *jsonable) ToMap() map[string]any {
	return s.value
}

func NewJSONable(value map[string]any) *jsonable {
	return &jsonable{
		value: value,
	}
}
