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
}
