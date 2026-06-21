package valueobject

import (
	"regexp"

	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
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
		return nil, errorc.NewDomain("name")
	}
	return &name{value}, nil
}

func (n *name) Value() string {
	return n.value
}
