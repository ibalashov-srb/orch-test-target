package orchtest

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"both_positive", 3, 4, 12},
		{"both_negative", -3, -4, 12},
		{"mixed_sign", 3, -4, -12},
		{"zero_times_value", 0, 7, 0},
		{"value_times_zero", 5, 0, 0},
		{"identity", 1, 99, 99},
		{"large_values", 1000, 999, 999000},
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
