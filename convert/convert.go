// The Convert package provides generalized type conversion utilities.
package convert

import (
	"errors"
	"strconv"
)

// StringSliceToIntSlice accepts a slice of strings and returns a slice of parsed ints.
//
// If any of the strings cannot be parsed to an integer, an error will be returned.
func StringSliceToIntSlice(strs []string) ([]int, error) {
	ints := make([]int, len(strs), len(strs))

	for i, str := range strs {
		anInt, err := strconv.ParseInt(str, 10, 0)
		if err != nil {
			return nil, errors.New("Failed to parse UserId[" + str + "] due to: " + err.Error())
		}

		ints[i] = int(anInt)
	}

	return ints, nil
}

// IntSliceToStringSlice converts a slice of ints to a slice of strings.
func IntSliceToStringSlice(ints []int) ([]string) {
	strings := make([]string, len(ints), len(ints))

	for i, val := range ints {
		strings[i] = strconv.Itoa(val)
	}

	return strings
}