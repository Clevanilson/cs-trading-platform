package entity_test

import (
	"testing"

	"github.com/clevanilson/cs-trading-platform/account_service/internal/domain/entity"
	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
)

func TestAccount(t *testing.T) {

	setup := func() {
		var err error
		pkgassert.Equals(t, err, nil)
	}

	t.Run("With valid data", func(t *testing.T) {
		setup()
		ID := "uuid"
		sut, err := entity.NewAccount(entity.AccountBuilder{
			Name: "Renoir",
			ID:   &ID,
		})
		pkgassert.NotEquals(t, sut, nil)
		pkgassert.Equals(t, err, nil)
		pkgassert.Equals(t, sut.Name(), "Renoir")
		pkgassert.Equals(t, sut.ID(), ID)
	})

	t.Run("With invalid name", func(t *testing.T) {
		setup()
		sut, err := entity.NewAccount(entity.AccountBuilder{
			Name: "Reno1r",
		})
		pkgassert.Equals(t, sut, nil)
		pkgassert.NotEquals(t, err, nil)
		pkgassert.Equals(t, err.Error(), "Invalid name")
	})
}
