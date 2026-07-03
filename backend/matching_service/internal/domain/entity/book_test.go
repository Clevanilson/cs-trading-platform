package entity_test

import (
	"testing"

	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
	pkgentity "github.com/clevanilson/cs-trading-platform/devpack/pkg/domain/entity"
	"github.com/clevanilson/cs-trading-platform/matching_service/internal/domain/entity"
)

func TestBook(t *testing.T) {
	t.Run("Insert buy order", func(t *testing.T) {
		book := entity.NewBook("BTC-USD")
		order1, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: "123",
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     10000,
			Amount:    1,
		})
		order2, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: "123",
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     100001,
			Amount:    1,
		})
		book.Insert(order1)
		book.Insert(order2)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, book.GetBestBuyOrder().ID(), order2.ID())
	})

	t.Run("Insert sell order", func(t *testing.T) {
		book := entity.NewBook("BTC-USD")
		order1, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: "123",
			MarketID:  "BTC-USD",
			Side:      "sell",
			Price:     10000,
			Amount:    1,
		})
		order2, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: "123",
			MarketID:  "BTC-USD",
			Side:      "sell",
			Price:     9000,
			Amount:    1,
		})
		book.Insert(order1)
		book.Insert(order2)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, book.GetBestSellOrder().ID(), order2.ID())
	})

	t.Run("Execute buy order", func(t *testing.T) {
		book := entity.NewBook("BTC-USD")
		order1, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: "123",
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     10000,
			Amount:    1,
		})
		book.Insert(order1)
		err = book.Execute(order1)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, order1.Amount(), 0)
	})

	t.Run("Execute sell order", func(t *testing.T) {
		book := entity.NewBook("BTC-USD")
		order1, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: "123",
			MarketID:  "BTC-USD",
			Side:      "sell",
			Price:     10000,
			Amount:    1,
		})
		book.Insert(order1)
		err = book.Execute(order1)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, order1.Amount(), 0)
	})
}
