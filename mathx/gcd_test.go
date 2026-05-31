package mathx

import "testing"

func TestGCD(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{0, 0, 0},
		{12, 8, 4},
		{-12, -8, 4},
		{-12, 8, 4},
		{7, 13, 1},
		{5, 0, 5},
		{0, 5, 5},
	}
	for _, c := range cases {
		if got := GCD(c.a, c.b); got != c.want {
			t.Errorf("GCD(%d, %d) = %d; want %d", c.a, c.b, got, c.want)
		}
	}
}
