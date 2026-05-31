package orchtest

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"zero times anything", 0, 5, 0},
		{"anything times zero", 7, 0, 0},
		{"positive * positive", 3, 4, 12},
		{"negative * positive", -2, 5, -10},
		{"negative * negative", -3, -4, 12},
		{"identity left", 1, 9, 9},
		{"identity right", 9, 1, 9},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := Multiply(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("Multiply(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}
