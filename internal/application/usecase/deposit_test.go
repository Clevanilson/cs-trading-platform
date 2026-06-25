package usecase_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	"github.com/clevanilson/cs-trading-platform/internal/domain/entity"
	infrarepository "github.com/clevanilson/cs-trading-platform/internal/infra/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
)

func TestDeposit(t *testing.T) {
	var repository repository.AccountRepository
	var sut usecase.Deposit
	var account entity.Account

	setup := func() {
		var err error
		repository = infrarepository.NewAccountMemoryRepository()
		sut = usecase.NewDeposit(repository)
		account, err = entity.NewAccount(entity.AccountBuilder{Name: "Dante"})
		err = repository.Save(account)
		assert.Equals(t, err, nil)
	}

	t.Run("With valid value", func(t *testing.T) {
		setup()
		err := sut.Execute(usecase.DepositInput{
			AccountID: account.ID(),
			AssetID:   "USD",
			Amount:    1000,
		})
		account, err = repository.GetByID(account.ID())
		assert.Equals(t, err, nil)
		assert.Equals(t, len(account.Balances()), 1)
		assert.Equals(t, account.Balances()[0].AssetID(), "USD")
		assert.Equals(t, account.Balances()[0].Amount(), 1000)
	})

	t.Run("With invalid value", func(t *testing.T) {
		err := sut.Execute(usecase.DepositInput{
			AccountID: account.ID(),
			AssetID:   "USD",
			Amount:    -1000,
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid amount")
	})

	t.Run("With a non-existent account", func(t *testing.T) {
		err := sut.Execute(usecase.DepositInput{
			AccountID: "cbce8b3e-c5fb-4118-87dc-db0897241c48",
			AssetID:   "USD",
			Amount:    1000,
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Account not found")
	})

}
