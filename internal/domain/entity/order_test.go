package entity_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/internal/domain/entity"
	valueobject "github.com/clevanilson/cs-trading-platform/internal/domain/value_object"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
)

func TestOrder(t *testing.T) {
	t.Run("With valid data", func(t *testing.T) {
		accountId := valueobject.NewID(nil).Value()
		sut, err := entity.NewOrder(entity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     1000,
			Amount:    10,
		})
		assert.Equals(t, err, nil)
		assert.NotEquals(t, sut, nil)
		assert.Equals(t, sut.AccountID(), accountId)
		assert.Equals(t, sut.MarketID(), "BTC-USD")
		assert.Equals(t, sut.MainAsset(), "BTC")
		assert.Equals(t, sut.PaymentAsset(), "USD")
		assert.Equals(t, sut.Side(), "buy")
		assert.Equals(t, sut.Price(), 1000)
		assert.Equals(t, sut.Amount(), 10)
		assert.Equals(t, sut.Status(), "open")
		assert.Equals(t, sut.CreatedAt().IsZero(), false)
	})

	t.Run("With invalid price", func(t *testing.T) {
		accountId := valueobject.NewID(nil).Value()
		sut, err := entity.NewOrder(entity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     -1000,
			Amount:    10,
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, sut, nil)
		assert.Equals(t, err.Error(), "Invalid price")
	})

	t.Run("With invalid amount", func(t *testing.T) {
		accountId := valueobject.NewID(nil).Value()
		sut, err := entity.NewOrder(entity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-USD",
			Side:      "sell",
			Price:     1000,
			Amount:    -10,
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, sut, nil)
		assert.Equals(t, err.Error(), "Invalid amount")
	})

	t.Run("With invalid side", func(t *testing.T) {
		accountId := valueobject.NewID(nil).Value()
		sut, err := entity.NewOrder(entity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-USD",
			Side:      "-",
			Price:     1000,
			Amount:    10,
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, sut, nil)
		assert.Equals(t, err.Error(), "Invalid side")
	})

	t.Run("With invalid accountId", func(t *testing.T) {
		sut, err := entity.NewOrder(entity.OrderBuilder{
			AccountID: "",
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     1000,
			Amount:    10,
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, sut, nil)
		assert.Equals(t, err.Error(), "Invalid accountID")
	})

	t.Run("With invalid market id", func(t *testing.T) {
		accountId := valueobject.NewID(nil).Value()
		sut, err := entity.NewOrder(entity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-UD",
			Side:      "sell",
			Price:     1000,
			Amount:    10,
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, sut, nil)
		assert.Equals(t, err.Error(), "Invalid marketID")
	})

	t.Run("With invalid status", func(t *testing.T) {
		accountId := valueobject.NewID(nil).Value()
		sut, err := entity.NewOrder(entity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     1000,
			Amount:    10,
			Status:    "invalid",
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, sut, nil)
		assert.Equals(t, err.Error(), "Invalid status")
	})

}
