package contains

import "testing"

func TestInt(t *testing.T) {
	// Negative cases
	if Int(0, []int{}) {
		t.Fatal("Expected value not to be in empty slice")
	} else if Int(0, []int{1, 2, 3}) {
		t.Fatal("Expected value not to be in slice")
	}

	// Positive Cases
	if !Int(0, []int{1, 2, 3, 0}) {
		t.Fatal("Expected value to be in slice")
	}
}