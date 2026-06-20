package valueobject_test

import (
	"testing"

	valueobject "github.com/clevanilson/cs-trading-platform/internal/domain/value_object"
	"github.com/clevanilson/cs-trading-platform/pkg/assert"
	"github.com/google/uuid"
)

func TestID(t *testing.T) {
	t.Run("Creating with nil value", func(t *testing.T) {
		sut := valueobject.NewID(nil)
		assert.Equals(t, uuid.Validate(sut.Value()), nil)
	})

	t.Run("Creating with uuid", func(t *testing.T) {
		value := "f9303a45-17f2-40cc-82f7-c6ef5f96a4c"
		sut := valueobject.NewID(&value)
		assert.Equals(t, sut.Value(), value)
	})
}
