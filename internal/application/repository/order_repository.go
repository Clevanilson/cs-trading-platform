package repository

import "github.com/clevanilson/cs-trading-platform/internal/domain/entity"

type OrderRepository interface {
	Save(order entity.Order) error
	Update(order entity.Order) error
	GetById(id string) (entity.Order, error)
}
