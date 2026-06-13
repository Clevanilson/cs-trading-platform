package assert

import "testing"

func Equals[TValue comparable](t *testing.T, value1, value2 TValue) {
	if value1 != value2 {
		t.Fatalf("🔴 Expected value: %v to equal %v", value1, value2)
	}
}
