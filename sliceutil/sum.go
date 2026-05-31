// Package sliceutil provides utility functions for slices.
package sliceutil

// Sum returns the arithmetic total of xs. A nil or empty slice yields 0.
func Sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}
