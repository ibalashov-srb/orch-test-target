package sliceutil

// Contains reports whether v is present in xs.
// A nil or empty slice yields false.
func Contains(xs []int, v int) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}
