package entity

import valueobject "github.com/clevanilson/cs-trading-platform/internal/domain/value_object"

type Account interface {
	Name() string
	ID() *int
}

type account struct {
	name valueobject.Name
	id   *int
}

type AccountBuilder struct {
	Name string
	id   *int
}

func NewAccount(builder AccountBuilder) (*account, error) {
	name, err := valueobject.NewName(builder.Name)
	if err != nil {
		return nil, err
	}
	return &account{name: name}, nil
}

func (a *account) Name() string {
	return a.name.Value()
}

func (a *account) ID() *int {
	return a.id
}
