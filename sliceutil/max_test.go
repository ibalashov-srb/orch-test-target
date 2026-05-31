package sliceutil

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		name    string
		xs      []int
		wantMax int
		wantOK  bool
	}{
		{"empty slice", []int{}, 0, false},
		{"nil slice", nil, 0, false},
		{"single element", []int{42}, 42, true},
		{"ascending", []int{1, 2, 3, 4, 5}, 5, true},
		{"descending", []int{5, 4, 3, 2, 1}, 5, true},
		{"all negative", []int{-3, -1, -2}, -1, true},
		{"mixed", []int{-5, 10, -3, 8}, 10, true},
		{"duplicates", []int{3, 3, 3}, 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotOK := Max(tt.xs)
			if gotOK != tt.wantOK {
				t.Errorf("Max(%v) ok = %v; want %v", tt.xs, gotOK, tt.wantOK)
			}
			if gotMax != tt.wantMax {
				t.Errorf("Max(%v) = %d; want %d", tt.xs, gotMax, tt.wantMax)
			}
		})
	}
}
