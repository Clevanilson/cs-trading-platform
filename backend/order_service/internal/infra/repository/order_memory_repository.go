package infrarepository

import (
	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"
)

func NewOrderMemoryRepository() *orderMemoryRepository {
	return &orderMemoryRepository{
		orders: make(map[string]entity.Order),
	}
}

type orderMemoryRepository struct {
	orders map[string]entity.Order
}

func (r *orderMemoryRepository) Save(order entity.Order) error {
	r.orders[order.ID()] = order
	return nil
}

func (r *orderMemoryRepository) GetByID(id string) (entity.Order, error) {
	order, ok := r.orders[id]
	if !ok {
		return nil, nil
	}
	return order, nil
}

func (r *orderMemoryRepository) Update(order entity.Order) error {
	_, ok := r.orders[order.ID()]
	if !ok {
		return pkgerror.NewNotFound("order")
	}
	r.orders[order.ID()] = order
	return nil
}
