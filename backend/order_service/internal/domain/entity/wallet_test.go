package entity_test

import (
	"testing"

	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
	"github.com/clevanilson/cs-trading-platform/order_service/internal/domain/entity"
)

func TestAccount(t *testing.T) {
	var sut entity.Wallet

	setup := func() {
		var err error
		sut, err = entity.NewWallet(entity.WalletBuilder{AccountID: "fe2aeda4-f5b1-40c2-b5d5-cb67dc66f5af"})
		pkgassert.Equals(t, err, nil)
	}

	t.Run("Deposit", func(t *testing.T) {
		t.Run("With valid amount and new asset", func(t *testing.T) {
			setup()
			err := sut.Deposit("USD", 100)
			err = sut.Deposit("BTC", 150)
			usd, err := sut.GetBalanceByAssetID("USD")
			btc, err := sut.GetBalanceByAssetID("BTC")
			pkgassert.Equals(t, err, nil)
			pkgassert.Equals(t, len(sut.Balances()), 2)
			pkgassert.Equals(t, usd.AssetID(), "USD")
			pkgassert.Equals(t, usd.Amount(), 100)
			pkgassert.Equals(t, btc.AssetID(), "BTC")
			pkgassert.Equals(t, btc.Amount(), 150)
		})

		t.Run("With valid amount and existing asset", func(t *testing.T) {
			setup()
			err := sut.Deposit("USD", 100)
			err = sut.Deposit("USD", 100)
			err = sut.Deposit("BTC", 150)
			err = sut.Deposit("BTC", 150)
			usd, err := sut.GetBalanceByAssetID("USD")
			btc, err := sut.GetBalanceByAssetID("BTC")
			pkgassert.Equals(t, err, nil)
			pkgassert.Equals(t, len(sut.Balances()), 2)
			pkgassert.Equals(t, usd.AssetID(), "USD")
			pkgassert.Equals(t, usd.Amount(), 200)
			pkgassert.Equals(t, btc.AssetID(), "BTC")
			pkgassert.Equals(t, btc.Amount(), 300)
		})

		t.Run("With invalid amount and new asset", func(t *testing.T) {
			setup()
			err := sut.Deposit("USD", -100)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Invalid amount")
		})

		t.Run("With invalid amount and existing asset", func(t *testing.T) {
			setup()
			err := sut.Deposit("USD", 100)
			err = sut.Deposit("USD", -100)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Invalid amount")
		})
	})

	t.Run("Withdraw", func(t *testing.T) {
		t.Run("With valid amount and existig asset", func(t *testing.T) {
			setup()
			err := sut.Deposit("USD", 100)
			err = sut.Deposit("BTC", 100)
			err = sut.Withdraw("USD", 50)
			err = sut.Withdraw("BTC", 100)
			usd, err := sut.GetBalanceByAssetID("USD")
			pkgassert.Equals(t, err, nil)
			btc, err := sut.GetBalanceByAssetID("BTC")
			pkgassert.Equals(t, err, nil)
			pkgassert.NotEquals(t, btc, nil)
			pkgassert.Equals(t, btc.Amount(), 0)
			pkgassert.Equals(t, len(sut.Balances()), 2)
			pkgassert.Equals(t, usd.AssetID(), "USD")
			pkgassert.Equals(t, usd.Amount(), 50)
			pkgassert.Equals(t, usd.Amount(), 50)
		})

		t.Run("With valid amount and non-existent asset", func(t *testing.T) {
			setup()
			err := sut.Withdraw("USD", 50)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Balance not found")
			usd, err := sut.GetBalanceByAssetID("USD")
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Balance not found")
			pkgassert.Equals(t, usd, nil)
		})

		t.Run("With invalid amount and exisiting asset", func(t *testing.T) {
			setup()
			err := sut.Deposit("BTC", 100)
			err = sut.Withdraw("BTC", 150)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Invalid amount")
			err = sut.Withdraw("321", 1)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Balance not found")
		})
	})

	t.Run("LockAmount", func(t *testing.T) {
		t.Run("With funds", func(t *testing.T) {
			setup()
			err := sut.Deposit("USD", 100)
			err = sut.Deposit("BTC", 100)
			order1, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.AccountID(),
				MarketID:  "BTC-USD",
				Side:      "buy",
				Price:     10,
				Amount:    5,
			})
			order2, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.AccountID(),
				MarketID:  "BTC-USD",
				Side:      "sell",
				Price:     10,
				Amount:    50,
			})
			err = sut.LockAmount(order1)
			asset, err := sut.GetBalanceByAssetID("USD")
			pkgassert.Equals(t, err, nil)
			pkgassert.Equals(t, asset.Amount(), 50)
			err = sut.LockAmount(order2)
			asset, err = sut.GetBalanceByAssetID("BTC")
			pkgassert.Equals(t, err, nil)
			pkgassert.Equals(t, asset.Amount(), 50)
		})

		t.Run("Without funds", func(t *testing.T) {
			setup()
			err := sut.Deposit("USD", 10)
			err = sut.Deposit("BTC", 10)
			order1, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.AccountID(),
				MarketID:  "BTC-USD",
				Side:      "buy",
				Price:     10,
				Amount:    5,
			})
			order2, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.AccountID(),
				MarketID:  "BTC-USD",
				Side:      "sell",
				Price:     10,
				Amount:    50,
			})
			err = sut.LockAmount(order1)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Insufficient funds")
			err = sut.LockAmount(order2)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Insufficient funds")
		})
	})

	t.Run("UnlockAmount", func(t *testing.T) {
		t.Run("With valid amount", func(t *testing.T) {
			setup()
			err := sut.Deposit("USD", 10)
			order1, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.AccountID(),
				MarketID:  "BTC-USD",
				Side:      "buy",
				Price:     10,
				Amount:    1,
			})
			err = sut.LockAmount(order1)
			err = sut.UnlockAmount(order1)
			asset, err := sut.GetBalanceByAssetID("USD")
			pkgassert.Equals(t, err, nil)
			pkgassert.Equals(t, asset.Amount(), 10)
		})

		t.Run("With invalid amount", func(t *testing.T) {
			setup()
			err := sut.Deposit("USD", 10)
			order1, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.AccountID(),
				MarketID:  "BTC-USD",
				Side:      "buy",
				Price:     10,
				Amount:    1,
			})
			err = sut.UnlockAmount(order1)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Invalid amount")
		})
	})

}
