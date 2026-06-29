package usecase_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/account_service/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/account_service/internal/application/usecase"
	infrarepository "github.com/clevanilson/cs-trading-platform/account_service/internal/infra/repository"
	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
)

func TestCreateAccount(t *testing.T) {
	var repository repository.AccountRepository
	var sut usecase.CreateAccount

	setup := func() {
		var err error
		repository = infrarepository.NewAccountMemoryRepository()
		sut = usecase.NewCreateAccount(repository)
		pkgassert.Equals(t, err, nil)
	}

	t.Run("With valid data", func(t *testing.T) {
		setup()
		output, err := sut.Execute(usecase.CreateAccountInput{
			Name: "Spider Man",
		})
		pkgassert.NotEquals(t, output, nil)
		pkgassert.Equals(t, err, nil)
	})

	t.Run("With invalid name", func(t *testing.T) {
		setup()
		output, err := sut.Execute(usecase.CreateAccountInput{
			Name: "Spider_Man",
		})
		pkgassert.Equals(t, output, nil)
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, err.Error(), "Invalid name")
	})
}
