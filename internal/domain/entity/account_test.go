package entity_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/internal/domain/entity"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
)

func TestAccount(t *testing.T) {

	t.Run("With valid data", func(t *testing.T) {
		ID := "uuid"
		sut, err := entity.NewAccount(entity.AccountBuilder{
			Name: "Renoir",
			ID:   &ID,
		})
		assert.NotEquals(t, sut, nil)
		assert.Equals(t, err, nil)
		assert.Equals(t, sut.Name(), "Renoir")
		assert.Equals(t, sut.ID(), ID)
	})

	t.Run("With invalid name", func(t *testing.T) {
		sut, err := entity.NewAccount(entity.AccountBuilder{
			Name: "Reno1r",
		})
		assert.Equals(t, sut, nil)
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid name")
	})

	t.Run("Deposit", func(t *testing.T) {
		t.Run("With valid amount and new asset", func(t *testing.T) {
			sut, err := entity.NewAccount(entity.AccountBuilder{
				Name: "Renoir",
			})
			err = sut.Deposit("USD", 100)
			err = sut.Deposit("BTC", 150)
			usd, err := sut.GetBalanceByAssetID("USD")
			btc, err := sut.GetBalanceByAssetID("BTC")
			assert.Equals(t, err, nil)
			assert.Equals(t, len(sut.Balances()), 2)
			assert.Equals(t, usd.AssetID(), "USD")
			assert.Equals(t, usd.Amount(), 100)
			assert.Equals(t, btc.AssetID(), "BTC")
			assert.Equals(t, btc.Amount(), 150)
		})

		t.Run("With valid amount and existing asset", func(t *testing.T) {
			sut, err := entity.NewAccount(entity.AccountBuilder{
				Name: "Renoir",
			})
			err = sut.Deposit("USD", 100)
			err = sut.Deposit("USD", 100)
			err = sut.Deposit("BTC", 150)
			err = sut.Deposit("BTC", 150)
			usd, err := sut.GetBalanceByAssetID("USD")
			btc, err := sut.GetBalanceByAssetID("BTC")
			assert.Equals(t, err, nil)
			assert.Equals(t, len(sut.Balances()), 2)
			assert.Equals(t, usd.AssetID(), "USD")
			assert.Equals(t, usd.Amount(), 200)
			assert.Equals(t, btc.AssetID(), "BTC")
			assert.Equals(t, btc.Amount(), 300)
		})

		t.Run("With invalid amount and new asset", func(t *testing.T) {
			sut, err := entity.NewAccount(entity.AccountBuilder{
				Name: "Renoir",
			})
			err = sut.Deposit("USD", -100)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Invalid amount")
		})

		t.Run("With invalid amount and existing asset", func(t *testing.T) {
			sut, err := entity.NewAccount(entity.AccountBuilder{
				Name: "Renoir",
			})
			err = sut.Deposit("USD", 100)
			err = sut.Deposit("USD", -100)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Invalid amount")
		})
	})

	t.Run("Withdraw", func(t *testing.T) {
		t.Run("With valid amount and existig asset", func(t *testing.T) {
			sut, err := entity.NewAccount(entity.AccountBuilder{
				Name: "Renoir",
			})
			err = sut.Deposit("USD", 100)
			err = sut.Deposit("BTC", 100)
			err = sut.Withdraw("USD", 50)
			err = sut.Withdraw("BTC", 100)
			usd, err := sut.GetBalanceByAssetID("USD")
			assert.Equals(t, err, nil)
			btc, err := sut.GetBalanceByAssetID("BTC")
			assert.Equals(t, err, nil)
			assert.NotEquals(t, btc, nil)
			assert.Equals(t, btc.Amount(), 0)
			assert.Equals(t, len(sut.Balances()), 2)
			assert.Equals(t, usd.AssetID(), "USD")
			assert.Equals(t, usd.Amount(), 50)
			assert.Equals(t, usd.Amount(), 50)
		})

		t.Run("With valid amount and non-existent asset", func(t *testing.T) {
			sut, err := entity.NewAccount(entity.AccountBuilder{
				Name: "Renoir",
			})
			err = sut.Withdraw("USD", 50)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Balance not found")
			usd, err := sut.GetBalanceByAssetID("USD")
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Balance not found")
			assert.Equals(t, usd, nil)
		})

		t.Run("With invalid amount and exisiting asset", func(t *testing.T) {
			sut, err := entity.NewAccount(entity.AccountBuilder{
				Name: "Verso",
			})
			err = sut.Deposit("BTC", 100)
			err = sut.Withdraw("BTC", 150)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Invalid amount")
			err = sut.Withdraw("321", 1)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Balance not found")
		})
	})

	t.Run("LockAmount", func(t *testing.T) {
		t.Run("With funds", func(t *testing.T) {
			sut, err := entity.NewAccount(entity.AccountBuilder{
				Name: "Verso",
			})
			err = sut.Deposit("USD", 100)
			err = sut.Deposit("BTC", 100)
			order1, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.ID(),
				MarketID:  "BTC-USD",
				Side:      "buy",
				Price:     10,
				Amount:    5,
			})
			order2, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.ID(),
				MarketID:  "BTC-USD",
				Side:      "sell",
				Price:     10,
				Amount:    50,
			})
			err = sut.LockAmount(order1)
			asset, err := sut.GetBalanceByAssetID("USD")
			assert.Equals(t, err, nil)
			assert.Equals(t, asset.Amount(), 50)
			err = sut.LockAmount(order2)
			asset, err = sut.GetBalanceByAssetID("BTC")
			assert.Equals(t, err, nil)
			assert.Equals(t, asset.Amount(), 50)
		})

		t.Run("Without funds", func(t *testing.T) {
			sut, err := entity.NewAccount(entity.AccountBuilder{
				Name: "Verso",
			})
			err = sut.Deposit("USD", 10)
			err = sut.Deposit("BTC", 10)
			order1, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.ID(),
				MarketID:  "BTC-USD",
				Side:      "buy",
				Price:     10,
				Amount:    5,
			})
			order2, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.ID(),
				MarketID:  "BTC-USD",
				Side:      "sell",
				Price:     10,
				Amount:    50,
			})
			err = sut.LockAmount(order1)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Insufficient funds")
			err = sut.LockAmount(order2)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Insufficient funds")
		})
	})

	t.Run("UnlockAmount", func(t *testing.T) {
		t.Run("With valid amount", func(t *testing.T) {
			sut, err := entity.NewAccount(entity.AccountBuilder{
				Name: "Verso",
			})
			err = sut.Deposit("USD", 10)
			order1, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.ID(),
				MarketID:  "BTC-USD",
				Side:      "buy",
				Price:     10,
				Amount:    1,
			})
			err = sut.LockAmount(order1)
			err = sut.UnlockAmount(order1)
			asset, err := sut.GetBalanceByAssetID("USD")
			assert.Equals(t, err, nil)
			assert.Equals(t, asset.Amount(), 10)
		})

		t.Run("With invalid amount", func(t *testing.T) {
			sut, err := entity.NewAccount(entity.AccountBuilder{
				Name: "Verso",
			})
			err = sut.Deposit("USD", 10)
			order1, err := entity.NewOrder(entity.OrderBuilder{
				AccountID: sut.ID(),
				MarketID:  "BTC-USD",
				Side:      "buy",
				Price:     10,
				Amount:    1,
			})
			err = sut.UnlockAmount(order1)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Invalid amount")
		})
	})

}
