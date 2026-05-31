package sliceutil

// Contains reports whether v is present in xs.
func Contains(xs []int, v int) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}
