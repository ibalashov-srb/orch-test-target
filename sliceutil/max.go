package sliceutil

// Max returns the maximum element of xs and true.
// For an empty slice it returns (0, false).
func Max(xs []int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	max := xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > max {
			max = xs[i]
		}
	}
	return max, true
}
