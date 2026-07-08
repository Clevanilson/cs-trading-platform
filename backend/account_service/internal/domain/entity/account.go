package entity

import (
	valueobject "github.com/clevanilson/cs-trading-platform/account_service/internal/domain/value_object"
	pkgvalueobject "github.com/clevanilson/cs-trading-platform/devpack/pkg/domain/value_object"
)

type Account interface {
	Name() string
	ID() string
}

type account struct {
	name valueobject.Name
	id   pkgvalueobject.ID
}

type AccountBuilder struct {
	Name string
	ID   string
}

func NewAccount(builder AccountBuilder) (*account, error) {
	name, err := valueobject.NewName(builder.Name)
	if err != nil {
		return nil, err
	}
	return &account{
		name: name,
		id:   pkgvalueobject.NewID(builder.ID),
	}, nil
}

func (a *account) Name() string {
	return a.name.Value()
}

func (a *account) ID() string {
	return a.id.Value()
}
