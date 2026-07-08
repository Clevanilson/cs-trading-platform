package pkgentity_test

import (
	"testing"

	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
	pkgentity "github.com/clevanilson/cs-trading-platform/devpack/pkg/domain/entity"
	pkgvalueobject "github.com/clevanilson/cs-trading-platform/devpack/pkg/domain/value_object"
)

func TestOrder(t *testing.T) {
	var sut pkgentity.Order

	setup := func() {
		var err error
		sut, err = pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: "123",
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     1000,
			Amount:    10,
		})
		pkgassert.Equals(t, err, nil)
		pkgassert.NotEquals(t, sut, nil)
	}

	t.Run("With valid data", func(t *testing.T) {
		accountId := pkgvalueobject.NewID("").Value()
		sut, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     1000,
			Amount:    10,
		})
		pkgassert.Equals(t, err, nil)
		pkgassert.NotEquals(t, sut, nil)
		pkgassert.Equals(t, sut.AccountID(), accountId)
		pkgassert.Equals(t, sut.MarketID(), "BTC-USD")
		pkgassert.Equals(t, sut.MainAsset(), "BTC")
		pkgassert.Equals(t, sut.PaymentAsset(), "USD")
		pkgassert.Equals(t, sut.Side(), "buy")
		pkgassert.Equals(t, sut.Price(), 1000)
		pkgassert.Equals(t, sut.Amount(), 10)
		pkgassert.Equals(t, sut.Status(), "open")
		pkgassert.Equals(t, sut.CreatedAt().IsZero(), false)
	})

	t.Run("With invalid price", func(t *testing.T) {
		accountId := pkgvalueobject.NewID("").Value()
		sut, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     -1000,
			Amount:    10,
		})
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, sut, nil)
		pkgassert.Equals(t, err.Error(), "Invalid price")
	})

	t.Run("With invalid amount", func(t *testing.T) {
		accountId := pkgvalueobject.NewID("").Value()
		sut, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-USD",
			Side:      "sell",
			Price:     1000,
			Amount:    -10,
		})
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, sut, nil)
		pkgassert.Equals(t, err.Error(), "Invalid amount")
	})

	t.Run("With invalid side", func(t *testing.T) {
		accountId := pkgvalueobject.NewID("").Value()
		sut, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-USD",
			Side:      "-",
			Price:     1000,
			Amount:    10,
		})
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, sut, nil)
		pkgassert.Equals(t, err.Error(), "Invalid side")
	})

	t.Run("With invalid accountId", func(t *testing.T) {
		sut, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: "",
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     1000,
			Amount:    10,
		})
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, sut, nil)
		pkgassert.Equals(t, err.Error(), "Invalid accountID")
	})

	t.Run("With invalid market id", func(t *testing.T) {
		accountId := pkgvalueobject.NewID("").Value()
		sut, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-UD",
			Side:      "sell",
			Price:     1000,
			Amount:    10,
		})
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, sut, nil)
		pkgassert.Equals(t, err.Error(), "Invalid marketID")
	})

	t.Run("With invalid status", func(t *testing.T) {
		accountId := pkgvalueobject.NewID("").Value()
		sut, err := pkgentity.NewOrder(pkgentity.OrderBuilder{
			AccountID: accountId,
			MarketID:  "BTC-USD",
			Side:      "buy",
			Price:     1000,
			Amount:    10,
			Status:    "invalid",
		})
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, sut, nil)
		pkgassert.Equals(t, err.Error(), "Invalid status")
	})

	t.Run("Fill with valid data", func(t *testing.T) {
		setup()
		err := sut.Fill(1, 1000)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, sut.Amount(), 9)
	})

	t.Run("Fill with invalid amount", func(t *testing.T) {
		setup()
		err := sut.Fill(-1, 1000)
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, err.Error(), "Invalid amount")
	})

}
