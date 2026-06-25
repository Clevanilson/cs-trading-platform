package infrarepository

import (
	"github.com/clevanilson/cs-trading-platform/internal/domain/entity"
	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
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

func (r *orderMemoryRepository) GetById(id string) (entity.Order, error) {
	order, ok := r.orders[id]
	if !ok {
		return nil, nil
	}
	return order, nil
}

func (r *orderMemoryRepository) Update(order entity.Order) error {
	_, ok := r.orders[order.ID()]
	if !ok {
		return errorc.NewNotFound("order")
	}
	r.orders[order.ID()] = order
	return nil
}
