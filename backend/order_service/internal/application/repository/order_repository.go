package repository

import pkgentity "github.com/clevanilson/cs-trading-platform/devpack/pkg/domain/entity"

type OrderRepository interface {
	Save(order pkgentity.Order) error
	Update(order pkgentity.Order) error
	GetByID(id string) (pkgentity.Order, error)
}
