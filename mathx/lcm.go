package mathx

// LCM returns the least common multiple of a and b.
// It returns 0 if either input is 0. Negative inputs are handled correctly
// because GCD normalises them. The division is performed before multiplication
// to reduce the risk of integer overflow.
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a / GCD(a, b) * b
}
