package usecase_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	infrarepository "github.com/clevanilson/cs-trading-platform/internal/infra/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
)

func TestWithdraw(t *testing.T) {
	repository := infrarepository.NewAccountMemoryRepository()
	createAccount := usecase.NewCreateAccount(repository)
	getAccount := usecase.NewGetAccount(repository)
	deposit := usecase.NewDeposit(repository)
	sut := usecase.NewWithdraw(repository)

	t.Run("With valid value", func(t *testing.T) {
		outputCreateAccount, err := createAccount.Execute(usecase.CreateAccountInput{
			Name: "Dante",
		})
		err = deposit.Execute(usecase.DepositInput{
			AccountID: outputCreateAccount.ID,
			AssetID:   "USD",
			Amount:    1000,
		})
		err = sut.Execute(usecase.WithdrawInput{
			AccountID: outputCreateAccount.ID,
			AssetID:   "USD",
			Amount:    500,
		})
		outputGetAccount, err := getAccount.Execute(usecase.GetAccountInput{
			ID: outputCreateAccount.ID,
		})
		assert.Equals(t, err, nil)
		assert.Equals(t, len(outputGetAccount.Balance), 1)
		assert.Equals(t, outputGetAccount.Balance[0].AssetID, "USD")
		assert.Equals(t, outputGetAccount.Balance[0].Amount, 500)
	})

	t.Run("With invalid value", func(t *testing.T) {
		outputCreateAccount, err := createAccount.Execute(usecase.CreateAccountInput{
			Name: "Dante",
		})
		err = deposit.Execute(usecase.DepositInput{
			AccountID: outputCreateAccount.ID,
			AssetID:   "USD",
			Amount:    1000,
		})
		err = sut.Execute(usecase.WithdrawInput{
			AccountID: outputCreateAccount.ID,
			AssetID:   "USD",
			Amount:    1001,
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid amount")
	})

}
