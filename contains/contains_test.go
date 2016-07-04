package contains

import "testing"

func TestInt(t *testing.T) {
	// Negative cases
	if contains := Int(0, []int{}); contains {
		t.Fatal("Expected value not to be in empty slice")
	} else if contains := Int(0, []int{1, 2, 3}); contains {
		t.Fatal("Expected value not to be in slice")
	}

	// Positive Cases
	if contains := Int(0, []int{1, 2, 3, 0}); !contains {
		t.Fatal("Expected value to be in slice")
	}
}