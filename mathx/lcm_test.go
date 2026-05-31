package mathx

import "testing"

func TestLCM(t *testing.T) {
	cases := []struct {
		a, b, want int
	}{
		{0, 5, 0},
		{4, 0, 0},
		{0, 0, 0},
		{4, 6, 12},
		{5, 5, 5},
		{-4, 6, 12},
	}
	for _, c := range cases {
		if got := LCM(c.a, c.b); got != c.want {
			t.Errorf("LCM(%d, %d) = %d; want %d", c.a, c.b, got, c.want)
		}
	}
}
