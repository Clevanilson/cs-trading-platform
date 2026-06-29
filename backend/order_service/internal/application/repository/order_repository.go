package repository

import "github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"

type OrderRepository interface {
	Save(order entity.Order) error
	Update(order entity.Order) error
	GetByID(id string) (entity.Order, error)
}
