package usecase_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	infrarepository "github.com/clevanilson/cs-trading-platform/internal/infra/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
)

func TestDeposit(t *testing.T) {
	repository := infrarepository.NewAccountMemoryRepository()
	createAccount := usecase.NewCreateAccount(repository)
	getAccount := usecase.NewGetAccount(repository)
	sut := usecase.NewDeposit(repository)

	t.Run("With valid value", func(t *testing.T) {
		outputCreateAccount, err := createAccount.Execute(usecase.CreateAccountInput{
			Name: "Dante",
		})
		err = sut.Execute(usecase.DepositInput{
			AccountID: outputCreateAccount.ID,
			AssetID:   "USD",
			Amount:    1000,
		})
		outputGetAccount, err := getAccount.Execute(usecase.GetAccountInput{
			ID: outputCreateAccount.ID,
		})
		assert.Equals(t, err, nil)
		assert.Equals(t, len(outputGetAccount.Balance), 1)
		assert.Equals(t, outputGetAccount.Balance[0].AssetID, "USD")
		assert.Equals(t, outputGetAccount.Balance[0].Amount, 1000)
	})

	t.Run("With invalid value", func(t *testing.T) {
		outputCreateAccount, err := createAccount.Execute(usecase.CreateAccountInput{
			Name: "Dante",
		})
		err = sut.Execute(usecase.DepositInput{
			AccountID: outputCreateAccount.ID,
			AssetID:   "USD",
			Amount:    -1000,
		})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid amount")
	})

}
