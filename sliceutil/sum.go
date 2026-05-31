package sliceutil

// Sum returns the arithmetic sum of xs. An empty or nil slice returns 0.
func Sum(xs []int) int {
	var total int
	for _, x := range xs {
		total += x
	}
	return total
}
