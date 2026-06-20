package usecase_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/internal/application/usecase"
	infrarepository "github.com/clevanilson/cs-trading-platform/internal/infra/repository"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
)

func TestCreateAccount(t *testing.T) {
	repository := infrarepository.NewAccountMemoryRepository()
	sut := usecase.NewCreateAccount(repository)

	t.Run("With valid data", func(t *testing.T) {
		output, err := sut.Execute(usecase.CreateAccountInput{
			Name: "Spider Man",
		})
		assert.NotEquals(t, output, nil)
		assert.Equals(t, err, nil)
	})

	t.Run("With invalid name", func(t *testing.T) {
		output, err := sut.Execute(usecase.CreateAccountInput{
			Name: "Spider_Man",
		})
		assert.Equals(t, output, nil)
		assert.NotEquals(t, err, nil)
		assert.Equals(t, err.Error(), "Invalid name")
	})
}
