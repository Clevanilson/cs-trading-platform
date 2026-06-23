package entity_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/internal/domain/entity"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
)

func TestBalance(t *testing.T) {
	t.Run("With valid data", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "123",
			Amount:  100,
		})
		assert.Equals(t, err, nil)
		assert.Equals(t, sut.AssetID(), "123")
		assert.Equals(t, sut.Amount(), 100)
	})

	t.Run("With invalid amount", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "123",
			Amount:  -12,
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid amount")
		assert.Equals(t, sut, nil)
	})

	t.Run("Deposit with valid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "123",
			Amount:  100,
		})
		assert.Equals(t, err, nil)
		err = sut.Deposit(50)
		assert.Equals(t, err, nil)
		assert.Equals(t, sut.Amount(), 150)
	})

	t.Run("Deposit with invalid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "123",
			Amount:  100,
		})
		assert.Equals(t, err, nil)
		err = sut.Deposit(-50)
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid amount")
	})

	t.Run("Withdraw with valid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "123",
			Amount:  100,
		})
		assert.Equals(t, err, nil)
		err = sut.Withdraw(50)
		assert.Equals(t, err, nil)
		assert.Equals(t, sut.Amount(), 50)
	})

	t.Run("Withdraw with invalid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "123",
			Amount:  100,
		})
		assert.Equals(t, err, nil)
		err = sut.Withdraw(-50)
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid amount")
		err = sut.Withdraw(150)
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid amount")
	})
}
