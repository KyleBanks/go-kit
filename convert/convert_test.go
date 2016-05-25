package convert

import (
	"strconv"
	"testing"
)

func TestStringSliceToIntSlice(t *testing.T) {
	// Test empty
	empty := make([]string, 0, 0)
	if res, err := StringSliceToIntSlice(empty); err != nil {
		t.Error(err)
	} else if len(res) > 0 {
		t.Error("Expected empty string slice to return empty int slice[]")
	}

	// Test proper use case
	proper := []string{"1", "-1", "0"}
	if res, err := StringSliceToIntSlice(proper); err != nil {
		t.Error(err)
	} else if len(res) != len(proper) {
		t.Error("Incorrect slice size returned. Got " + strconv.Itoa(len(res)) + " Expected " + strconv.Itoa(len(proper)))
	} else if res[0] != 1 || res[1] != -1 || res[2] != 0 {
		t.Error("Incorrect slice returned: ", res)
	}

	// Test an invalid element
	invalid := []string{"1", "invalid"}
	if _, err := StringSliceToIntSlice(invalid); err == nil {
		t.Error("Expected error!")
	}
}
