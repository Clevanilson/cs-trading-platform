package entity_test

import (
	"testing"

	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"
)

func TestBalance(t *testing.T) {
	var sut entity.Balance

	setup := func() {
		var err error
		sut, err = entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  100,
		})
		pkgassert.Equals(t, err, nil)
	}

	t.Run("With valid data", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  100,
		})
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, sut.AssetID(), "BTC")
		pkgassert.Equals(t, sut.Amount(), 100)
	})

	t.Run("With invalid amount", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  -12,
		})
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, err.Error(), "Invalid amount")
		pkgassert.Equals(t, sut, nil)
	})

	t.Run("Deposit with valid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  100,
		})
		pkgassert.Equals(t, err, nil)
		err = sut.Deposit(50)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, sut.Amount(), 150)
	})

	t.Run("Deposit with invalid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  100,
		})
		pkgassert.Equals(t, err, nil)
		err = sut.Deposit(-50)
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, err.Error(), "Invalid amount")
	})

	t.Run("Withdraw with valid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  100,
		})
		pkgassert.Equals(t, err, nil)
		err = sut.Withdraw(50)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, sut.Amount(), 50)
	})

	t.Run("Withdraw with invalid value", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "BTC",
			Amount:  100,
		})
		pkgassert.Equals(t, err, nil)
		err = sut.Withdraw(-50)
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, err.Error(), "Invalid amount")
		err = sut.Withdraw(150)
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, err.Error(), "Invalid amount")
	})

	t.Run("LockAmount", func(t *testing.T) {
		t.Run("With valid amount", func(t *testing.T) {
			setup()
			err := sut.LockAmount(50)
			pkgassert.NotEquals(t, sut, nil)
			err = sut.LockAmount(50)
			pkgassert.Equals(t, err, nil)
			pkgassert.Equals(t, sut.Amount(), 0)
		})

		t.Run("With invalid amount", func(t *testing.T) {
			setup()
			err := sut.LockAmount(150)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Insufficient funds")
			pkgassert.Equals(t, sut.Amount(), 100)
			err = sut.LockAmount(-50)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Invalid amount")
			pkgassert.Equals(t, sut.Amount(), 100)
		})
	})

	t.Run("UnlockAmount", func(t *testing.T) {
		t.Run("With valid amount", func(t *testing.T) {
			setup()
			err := sut.LockAmount(50)
			pkgassert.NotEquals(t, sut, nil)
			err = sut.UnlockAmount(25)
			pkgassert.Equals(t, err, nil)
			pkgassert.Equals(t, sut.Amount(), 75)
			err = sut.UnlockAmount(25)
			pkgassert.Equals(t, err, nil)
			pkgassert.Equals(t, sut.Amount(), 100)
		})

		t.Run("With invalid amount", func(t *testing.T) {
			setup()
			err := sut.LockAmount(50)
			err = sut.UnlockAmount(100)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Invalid amount")
			pkgassert.Equals(t, sut.Amount(), 50)
			err = sut.UnlockAmount(-50)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Invalid amount")
			pkgassert.Equals(t, sut.Amount(), 50)
		})
	})

	t.Run("With invalid asset ID", func(t *testing.T) {
		sut, err := entity.NewBalance(entity.BalanceBuilder{
			AssetID: "",
			Amount:  100,
		})
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, err.Error(), "Invalid asset ID")
		pkgassert.Equals(t, sut, nil)
	})
}
