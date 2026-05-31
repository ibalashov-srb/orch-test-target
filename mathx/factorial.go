package mathx

import "errors"

// Factorial returns n! for n >= 0. It returns an error for negative inputs.
func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial of negative number")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}
