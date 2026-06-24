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
			assert.NotEquals(t, err, nil)
			assert.Equals(t, btc, nil)
			assert.Equals(t, err.Error(), "Balance not found")
			assert.Equals(t, len(sut.Balances()), 1)
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
			err = sut.Deposit("123", 100)
			err = sut.Withdraw("123", 150)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Invalid amount")
			err = sut.Withdraw("321", 1)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), "Balance not found")
		})
	})

}
