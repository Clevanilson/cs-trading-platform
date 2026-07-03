package infrarepository

import (
	pkgentity "github.com/clevanilson/cs-trading-platform/devpack/pkg/domain/entity"
	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
)

func NewOrderMemoryRepository() *orderMemoryRepository {
	return &orderMemoryRepository{
		orders: make(map[string]pkgentity.Order),
	}
}

type orderMemoryRepository struct {
	orders map[string]pkgentity.Order
}

func (r *orderMemoryRepository) Save(order pkgentity.Order) error {
	r.orders[order.ID()] = order
	return nil
}

func (r *orderMemoryRepository) GetByID(id string) (pkgentity.Order, error) {
	order, ok := r.orders[id]
	if !ok {
		return nil, nil
	}
	return order, nil
}

func (r *orderMemoryRepository) Update(order pkgentity.Order) error {
	_, ok := r.orders[order.ID()]
	if !ok {
		return pkgerror.NewNotFound("order")
	}
	r.orders[order.ID()] = order
	return nil
}
