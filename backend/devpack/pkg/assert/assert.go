package pkgassert

import (
	"fmt"
	"testing"
)

func Equals[TValue comparable](t *testing.T, value1, value2 TValue) {
	if value1 != value2 {
		t.Fatalf("🔴 Expected value: %v to equal %v", value1, value2)
	}
}

func NotEquals[TValue comparable](t *testing.T, value1, value2 TValue) {
	if value1 == value2 {
		t.Fatalf("🔴 Expected value: %v to not equal %v", value1, value2)
	}
}

func Each[TValue comparable](t *testing.T, values []TValue, callback func(value TValue)) {
	for _, value := range values {
		t.Run(fmt.Sprintf("%v", value), func(_ *testing.T) {
			callback(value)
		})
	}

}
