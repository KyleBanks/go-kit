package contains

// Int returns true if the slice of ints contains the value provided.
func Int(val int, arr []int) bool {
	iarr := make([]interface{}, len(arr))
	for i, a := range arr {
		iarr[i] = a
	}

	return contains(val, iarr)
}

// Uint returns true if the slice of uints contains the value provided.
func Uint(val uint, arr []uint) bool {
	iarr := make([]interface{}, len(arr))
	for i, a := range arr {
		iarr[i] = a
	}

	return contains(val, iarr)
}

func contains(v interface{}, arr []interface{}) bool {
	for _, i := range arr {
		if i == v {
			return true
		}
	}

	return false
}