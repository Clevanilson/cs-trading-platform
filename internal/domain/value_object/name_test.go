package valueobject_test

import (
	"testing"

	valueobject "github.com/clevanilson/cs-trading-platform/internal/domain/value_object"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
	"github.com/clevanilson/cs-trading-platform/pkg/errorc"
)

func TestName(t *testing.T) {
	t.Run("With valid value", func(t *testing.T) {
		assert.Each(t, []string{"Maelle", "Eren Yeager", "Noctis Lucis Caelum"}, func(value string) {
			sut, err := valueobject.NewName(value)
			assert.Equals(t, err, nil)
			assert.NotEquals(t, sut, nil)
			assert.Equals(t, sut.Value(), value)
		})
	})

	t.Run("With invalid value", func(t *testing.T) {
		values := []string{
			"M@elle",
			"",
			"Noct1s Lucis Caelum",
			"John_Doe",
		}
		assert.Each(t, values, func(value string) {
			sut, err := valueobject.NewName(value)
			assert.NotEquals(t, err, nil)
			assert.Equals(t, err.Error(), errorc.NewDomain("name").Error())
			assert.Equals(t, sut, nil)
		})
	})
}
