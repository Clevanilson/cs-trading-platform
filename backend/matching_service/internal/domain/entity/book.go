package entity

import (
	"sort"

	pkgentity "github.com/clevanilson/cs-trading-platform/devpack/pkg/domain/entity"
	pkgerror "github.com/clevanilson/cs-trading-platform/devpack/pkg/error"
)

type Book interface {
	MarketID() string
	Insert(order pkgentity.Order)
	GetBestBuyOrder() pkgentity.Order
	GetBestSellOrder() pkgentity.Order
	Execute(order pkgentity.Order) error
}

type book struct {
	marketID   string
	buyOrders  []pkgentity.Order
	sellOrders []pkgentity.Order
}

func NewBook(marketID string) *book {
	return &book{
		marketID:   marketID,
		buyOrders:  make([]pkgentity.Order, 0),
		sellOrders: make([]pkgentity.Order, 0),
	}
}

func (b *book) MarketID() string {
	return b.marketID
}

func (b *book) Insert(order pkgentity.Order) {
	if order.Side() == "buy" {
		b.buyOrders = append(b.buyOrders, order)
		sort.Slice(b.buyOrders, func(i, j int) bool {
			return b.buyOrders[i].Price() > b.buyOrders[j].Price()
		})
	} else {
		b.sellOrders = append(b.sellOrders, order)
		sort.Slice(b.sellOrders, func(i, j int) bool {
			return b.sellOrders[i].Price() < b.sellOrders[j].Price()
		})
	}
}

func (b *book) GetBestBuyOrder() pkgentity.Order {
	return b.buyOrders[0]
}

func (b *book) GetBestSellOrder() pkgentity.Order {
	return b.sellOrders[0]
}

func (b *book) Execute(order pkgentity.Order) error {
	if order.Amount() == 0 {
		return nil
	}
	if order.Side() == "buy" {
		bestSellOrder := b.GetBestSellOrder()
		if bestSellOrder.Price() > order.Price() {
			return pkgerror.NewDomain("Best sell order is higher than buy order")
		}
	}
	if order.Side() == "sell" {
		bestBuyOrder := b.GetBestBuyOrder()
		if bestBuyOrder.Price() < order.Price() {
			return pkgerror.NewDomain("Best buy order is lower than sell order")
		}
	}
	return nil
}
