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
	sut := usecase.NewGetAccount(repository)

	t.Run("With existing account", func(t *testing.T) {
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

	})
}
