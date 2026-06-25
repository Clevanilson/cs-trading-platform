package entity_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/internal/domain/entity"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
)

func TestBalance(t *testing.T) {
	t.Run("With valid data", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  100,
		})
		assert.Equals(t, err, nil)
		assert.Equals(t, sut.AssetID(), "BTC")
		assert.Equals(t, sut.Amount(), 100)
	})

	t.Run("With invalid amount", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  -12,
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid amount")
		assert.Equals(t, sut, nil)
	})

	t.Run("Deposit with valid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  100,
		})
		assert.Equals(t, err, nil)
		err = sut.Deposit(50)
		assert.Equals(t, err, nil)
		assert.Equals(t, sut.Amount(), 150)
	})

	t.Run("Deposit with invalid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  100,
		})
		assert.Equals(t, err, nil)
		err = sut.Deposit(-50)
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid amount")
	})

	t.Run("Withdraw with valid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  100,
		})
		assert.Equals(t, err, nil)
		err = sut.Withdraw(50)
		assert.Equals(t, err, nil)
		assert.Equals(t, sut.Amount(), 50)
	})

	t.Run("Withdraw with invalid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
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

	t.Run("LockAmount", func(t *testing.T) {
		t.Run("With valid amount", func(t *testing.T) {
			sut, err := entity.NewBalance(entity.BalanceBuilder{
				AssetID: "BTC",
				Amount:  100,
			})
			assert.NotEquals(t, sut, nil)
			err = sut.LockAmount(50)
			err = sut.LockAmount(50)
			assert.Equals(t, err, nil)
			assert.Equals(t, sut.Amount(), 0)
		})

		t.Run("With invalid amount", func(t *testing.T) {
			sut, err := entity.NewBalance(entity.BalanceBuilder{
				AssetID: "BTC",
				Amount:  100,
			})
			assert.NotEquals(t, sut, nil)
			err = sut.LockAmount(150)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Insufficient funds")
			assert.Equals(t, sut.Amount(), 100)
			err = sut.LockAmount(-50)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Invalid amount")
			assert.Equals(t, sut.Amount(), 100)
		})
	})

	t.Run("UnlockAmount", func(t *testing.T) {
		t.Run("With valid amount", func(t *testing.T) {
			sut, err := entity.NewBalance(entity.BalanceBuilder{
				AssetID: "BTC",
				Amount:  100,
			})
			assert.NotEquals(t, sut, nil)
			err = sut.LockAmount(50)
			err = sut.UnlockAmount(25)
			assert.Equals(t, err, nil)
			assert.Equals(t, sut.Amount(), 75)
			err = sut.UnlockAmount(25)
			assert.Equals(t, err, nil)
			assert.Equals(t, sut.Amount(), 100)
		})

		t.Run("With invalid amount", func(t *testing.T) {
			sut, err := entity.NewBalance(entity.BalanceBuilder{
				AssetID: "BTC",
				Amount:  100,
			})
			sut.LockAmount(50)
			err = sut.UnlockAmount(100)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Invalid amount")
			assert.Equals(t, sut.Amount(), 50)
			err = sut.UnlockAmount(-50)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Invalid amount")
			assert.Equals(t, sut.Amount(), 50)
		})
	})
}
