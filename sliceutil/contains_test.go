package sliceutil

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		v    int
		want bool
	}{
		{"empty slice", []int{}, 42, false},
		{"nil slice", nil, 42, false},
		{"single element match", []int{42}, 42, true},
		{"single element no match", []int{7}, 42, false},
		{"first element", []int{42, 1, 2, 3}, 42, true},
		{"last element", []int{1, 2, 3, 42}, 42, true},
		{"middle element", []int{1, 42, 3}, 42, true},
		{"absent value", []int{1, 2, 3, 4, 5}, 99, false},
		{"negative value present", []int{-3, -1, -2}, -1, true},
		{"negative value absent", []int{-3, -1, -2}, -4, false},
		{"duplicates", []int{5, 5, 5}, 5, true},
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
