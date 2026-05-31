package mathx

import "testing"

func TestIsPrime(t *testing.T) {
	cases := []struct {
		n    int
		want bool
	}{
		{-1, false},
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{97, true},
		{100, false},
	}
	for _, c := range cases {
		if got := IsPrime(c.n); got != c.want {
			t.Errorf("IsPrime(%d) = %v; want %v", c.n, got, c.want)
		}
	}
}
