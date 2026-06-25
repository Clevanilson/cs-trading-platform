package usecase_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	"github.com/clevanilson/cs-trading-platform/internal/domain/entity"
	infrarepository "github.com/clevanilson/cs-trading-platform/internal/infra/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
)

func TestGetAccoun(t *testing.T) {
	repository := infrarepository.NewAccountMemoryRepository()
	var sut usecase.GetAccount

	setup := func() {
		var err error
		repository = infrarepository.NewAccountMemoryRepository()
		sut = usecase.NewGetAccount(repository)
		assert.Equals(t, err, nil)
	}

	t.Run("With existing account", func(t *testing.T) {
		setup()
		ID := "uuid"
		account, err := entity.NewAccount(entity.AccountBuilder{
			Name: "Jhon Doe",
			ID:   &ID,
		})
		err = repository.Save(account)
		output, err := sut.Execute(usecase.GetAccountInput{ID})
		assert.Equals(t, err, nil)
		assert.Equals(t, output.Name, account.Name())
		assert.Equals(t, output.ID, account.ID())
		assert.Equals(t, len(output.Balance), 0)
	})

	t.Run("With a non-existent account", func(t *testing.T) {
		setup()
		output, err := sut.Execute(usecase.GetAccountInput{ID: "cbce8b3e-c5fb-4118-87dc-db0897241c48"})
		assert.NotEquals(t, err, nil)
		assert.Equals(t, output, nil)
		assert.Equals(t, err.Error(), "Account not found")
	})
}
