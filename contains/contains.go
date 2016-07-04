package contains

// Int returns true if the slice of ints contains the value provided.
func Int(val int, arr []int) bool {
	for _, i := range arr {
		if i == val {
			return true
		}
	}

	return false
}
