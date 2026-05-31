package mathx

import "testing"

func TestGCD(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{name: "both_zero", a: 0, b: 0, want: 0},
		{name: "zero_a", a: 0, b: 5, want: 5},
		{name: "zero_b", a: 5, b: 0, want: 5},
		{name: "positive", a: 12, b: 8, want: 4},
		{name: "negative_a", a: -12, b: 8, want: 4},
		{name: "negative_b", a: 12, b: -8, want: 4},
		{name: "both_negative", a: -12, b: -8, want: 4},
		{name: "coprime", a: 7, b: 13, want: 1},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := GCD(tc.a, tc.b); got != tc.want {
				t.Errorf("GCD(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestLCM(t *testing.T) {
	cases := []struct {
		name string
		a, b int
		want int
	}{
		{name: "zero_a", a: 0, b: 5, want: 0},
		{name: "zero_b", a: 4, b: 0, want: 0},
		{name: "both_zero", a: 0, b: 0, want: 0},
		{name: "positive", a: 4, b: 6, want: 12},
		{name: "same", a: 5, b: 5, want: 5},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := LCM(tc.a, tc.b); got != tc.want {
				t.Errorf("LCM(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want bool
	}{
		{name: "negative", n: -1, want: false},
		{name: "zero", n: 0, want: false},
		{name: "one", n: 1, want: false},
		{name: "two", n: 2, want: true},
		{name: "three", n: 3, want: true},
		{name: "four", n: 4, want: false},
		{name: "large_prime", n: 97, want: true},
		{name: "large_composite", n: 100, want: false},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := IsPrime(tc.n); got != tc.want {
				t.Errorf("IsPrime(%d) = %v; want %v", tc.n, got, tc.want)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	cases := []struct {
		name    string
		n       int
		want    int
		wantErr bool
	}{
		{name: "negative", n: -1, want: 0, wantErr: true},
		{name: "zero", n: 0, want: 1, wantErr: false},
		{name: "one", n: 1, want: 1, wantErr: false},
		{name: "five", n: 5, want: 120, wantErr: false},
		{name: "ten", n: 10, want: 3628800, wantErr: false},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := Factorial(tc.n)
			if tc.wantErr {
				if err == nil {
					t.Errorf("Factorial(%d): expected error, got nil", tc.n)
				}
			} else {
				if err != nil {
					t.Errorf("Factorial(%d): unexpected error: %v", tc.n, err)
				}
				if got != tc.want {
					t.Errorf("Factorial(%d) = %d; want %d", tc.n, got, tc.want)
				}
			}
		})
	}
}
