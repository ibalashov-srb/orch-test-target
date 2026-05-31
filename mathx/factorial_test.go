package mathx

import "testing"

func TestFactorial(t *testing.T) {
	cases := []struct {
		n       int
		want    int
		wantErr bool
	}{
		{-1, 0, true},
		{0, 1, false},
		{1, 1, false},
		{5, 120, false},
		{10, 3628800, false},
	}
	for _, c := range cases {
		got, err := Factorial(c.n)
		if c.wantErr {
			if err == nil {
				t.Errorf("Factorial(%d): expected error, got nil", c.n)
			}
			if got != 0 {
				t.Errorf("Factorial(%d): expected 0 on error, got %d", c.n, got)
			}
		} else {
			if err != nil {
				t.Errorf("Factorial(%d): unexpected error: %v", c.n, err)
			}
			if got != c.want {
				t.Errorf("Factorial(%d) = %d, want %d", c.n, got, c.want)
			}
		}
	}
}
