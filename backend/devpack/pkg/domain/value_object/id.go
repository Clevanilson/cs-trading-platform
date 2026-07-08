package pkgvalueobject

import "github.com/google/uuid"

type ID interface {
	Value() string
}

type id struct {
	value string
}

func (v id) Value() string {
	return v.value
}

func NewID(value string) *id {
	var _value string
	if value != "" {
		_value = value
	} else {
		_value = uuid.New().String()
	}
	return &id{value: _value}
}
