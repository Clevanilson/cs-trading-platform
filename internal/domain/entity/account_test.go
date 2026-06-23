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

	t.Run("Deposit with valid value", func(t *testing.T) {
		sut, err := entity.NewAccount(entity.AccountBuilder{
			Name: "Renoir",
		})
		err = sut.Deposit("123", 100)
		assert.Equals(t, err, nil)
		balance, err := sut.GetBalanceByAssetID("123")
		assert.Equals(t, err, nil)
		assert.Equals(t, balance.Amount(), 100)
		err = sut.Deposit("123", 100)
		assert.Equals(t, balance.Amount(), 200)
	})

	t.Run("Deposit with invalid value", func(t *testing.T) {
		sut, err := entity.NewAccount(entity.AccountBuilder{
			Name: "Renoir",
		})
		err = sut.Deposit("123", -100)
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid amount")
	})

	t.Run("Withdraw with valid amount", func(t *testing.T) {
		sut, err := entity.NewAccount(entity.AccountBuilder{
			Name: "Renoir",
		})
		err = sut.Deposit("123", 100)
		err = sut.Withdraw("123", 50)
		balance, err := sut.GetBalanceByAssetID("123")
		assert.Equals(t, err, nil)
		assert.Equals(t, balance.Amount(), 50)
	})

	t.Run("Withdraw with invalid amount", func(t *testing.T) {
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
}
