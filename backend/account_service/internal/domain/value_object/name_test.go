package valueobject_test

import (
	"testing"

	valueobject "github.com/clevanilson/cs-trading-platform/account_service/internal/domain/value_object"
	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
)

func TestName(t *testing.T) {
	t.Run("With valid value", func(t *testing.T) {
		pkgassert.Each(t, []string{"Maelle", "Eren Yeager", "Noctis Lucis Caelum"}, func(value string) {
			sut, err := valueobject.NewName(value)
			pkgassert.Equals(t, err, nil)
			pkgassert.NotEquals(t, sut, nil)
			pkgassert.Equals(t, sut.Value(), value)
		})
	})

	t.Run("With invalid value", func(t *testing.T) {
		values := []string{
			"M@elle",
			"",
			"Noct1s Lucis Caelum",
			"John_Doe",
		}
		pkgassert.Each(t, values, func(value string) {
			sut, err := valueobject.NewName(value)
			pkgassert.NotEquals(t, err, nil)
			pkgassert.Equals(t, err.Error(), "Invalid name")
			pkgassert.Equals(t, sut, nil)
		})
	})
}
