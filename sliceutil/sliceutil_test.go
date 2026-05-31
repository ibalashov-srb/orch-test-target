package sliceutil

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		want int
	}{
		{"empty", []int{}, 0},
		{"all-positive", []int{1, 2, 3, 4, 5}, 15},
		{"mixed-signs", []int{-5, 10, -3, 8}, 10},
		{"single-element", []int{7}, 7},
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

func TestMax(t *testing.T) {
	tests := []struct {
		name    string
		xs      []int
		wantMax int
		wantOK  bool
	}{
		{"empty", []int{}, 0, false},
		{"single", []int{42}, 42, true},
		{"ascending", []int{1, 2, 3, 4, 5}, 5, true},
		{"descending", []int{5, 4, 3, 2, 1}, 5, true},
		{"all-negative", []int{-3, -1, -2}, -1, true},
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

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		v    int
		want bool
	}{
		{"present", []int{1, 2, 3}, 2, true},
		{"absent", []int{1, 2, 3}, 99, false},
		{"empty", []int{}, 1, false},
		{"first", []int{42, 1, 2, 3}, 42, true},
		{"last", []int{1, 2, 3, 42}, 42, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.xs, tt.v)
			if got != tt.want {
				t.Errorf("Contains(%v, %d) = %v; want %v", tt.xs, tt.v, got, tt.want)
			}
		})
	}
}
