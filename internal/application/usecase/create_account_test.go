package usecase_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/internal/application/repository"
	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	infrarepository "github.com/clevanilson/cs-trading-platform/internal/infra/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
)

func TestCreateAccount(t *testing.T) {
	var repository repository.AccountRepository
	var sut usecase.CreateAccount

	setup := func() {
		var err error
		repository = infrarepository.NewAccountMemoryRepository()
		sut = usecase.NewCreateAccount(repository)
		assert.Equals(t, err, nil)
	}

	t.Run("With valid data", func(t *testing.T) {
		setup()
		output, err := sut.Execute(usecase.CreateAccountInput{
			Name: "Spider Man",
		})
		assert.NotEquals(t, output, nil)
		assert.Equals(t, err, nil)
	})

	t.Run("With invalid name", func(t *testing.T) {
		setup()
		output, err := sut.Execute(usecase.CreateAccountInput{
			Name: "Spider_Man",
		})
		assert.Equals(t, output, nil)
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid name")
	})
}
