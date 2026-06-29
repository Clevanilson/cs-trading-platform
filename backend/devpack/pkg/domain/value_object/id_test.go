package pkgvalueobject_test

import (
	"testing"

	pkgassert "github.com/clevanilson/cs-trading-platform/devpack/pkg/assert"
	pkgvalueobject "github.com/clevanilson/cs-trading-platform/devpack/pkg/domain/value_object"
	"github.com/google/uuid"
)

func TestID(t *testing.T) {
	t.Run("Creating with nil value", func(t *testing.T) {
		sut := pkgvalueobject.NewID(nil)
		pkgassert.Equals(t, uuid.Validate(sut.Value()), nil)
	})

	t.Run("Creating with uuid", func(t *testing.T) {
		value := "f9303a45-17f2-40cc-82f7-c6ef5f96a4c"
		sut := pkgvalueobject.NewID(&value)
		pkgassert.Equals(t, sut.Value(), value)
	})
}
