package sliceutil

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		want int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single element", []int{7}, 7},
		{"all positive", []int{1, 2, 3, 4, 5}, 15},
		{"all negative", []int{-1, -2, -3}, -6},
		{"mixed", []int{-5, 10, -3, 8}, 10},
		{"zeros", []int{0, 0, 0}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.xs)
			if got != tt.want {
				t.Errorf("Sum(%v) = %d; want %d", tt.xs, got, tt.want)
			}
		})
	}
}
