package usecase_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/account_service/internal/application/usecase"
	"github.com/clevanilson/cs-trading-platform/account_service/internal/domain/entity"
	infrarepository "github.com/clevanilson/cs-trading-platform/account_service/internal/infra/repository"
	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
)

func TestGetAccoun(t *testing.T) {
	repository := infrarepository.NewAccountMemoryRepository()
	var sut usecase.GetAccount

	setup := func() {
		var err error
		repository = infrarepository.NewAccountMemoryRepository()
		sut = usecase.NewGetAccount(repository)
		pkgassert.Equals(t, err, nil)
	}

	t.Run("With existing account", func(t *testing.T) {
		setup()
		account, err := entity.NewAccount(entity.AccountBuilder{
			Name: "Jhon Doe",
			ID:   "uuid",
		})
		err = repository.Save(account)
		output, err := sut.Execute(usecase.GetAccountInput{ID: "uuid"})
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, output.Name, account.Name())
		pkgassert.Equals(t, output.ID, account.ID())
	})

	t.Run("With a non-existent account", func(t *testing.T) {
		setup()
		output, err := sut.Execute(usecase.GetAccountInput{ID: "cbce8b3e-c5fb-4118-87dc-db0897241c48"})
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, output, nil)
		pkgassert.Equals(t, err.Error(), "Account not found")
	})
}
