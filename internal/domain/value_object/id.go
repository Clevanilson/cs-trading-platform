package valueobject

import "github.com/google/uuid"

type ID struct {
	value string
}

func (id ID) Value() string {
	return id.value
}

func NewID(value *string) *ID {
	var id string
	if value != nil {
		id = *value
	} else {
		id = uuid.New().String()
	}
	return &ID{value: id}
}
