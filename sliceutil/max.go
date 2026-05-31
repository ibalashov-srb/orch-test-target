package sliceutil

// Max returns the largest element in xs and true.
// For a nil or empty slice it returns (0, false).
func Max(xs []int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	max := xs[0]
	for _, x := range xs[1:] {
		if x > max {
			max = x
		}
	}
	return max, true
}
