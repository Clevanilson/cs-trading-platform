package valueobject

import (
	"regexp"

	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
)

type Name interface {
	Value() string
}

type name struct {
	value string
}

func NewName(value string) (*name, error) {
	pattern := regexp.MustCompile(`^[a-zA-Z\s]{2,255}$`)
	isValid := pattern.MatchString(value)
	if !isValid {
		return nil, pkgerror.NewDomain("Invalid name")
	}
	return &name{value}, nil
}

func (n *name) Value() string {
	return n.value
}
