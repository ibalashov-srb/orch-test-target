package sliceutil

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		name string
		xs   []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{5}, 5},
		{"positive", []int{1, 2, 3}, 6},
		{"mixed_signs", []int{-1, -2, 3}, 0},
		{"duplicates", []int{2, 2, 2}, 6},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Sum(tc.xs); got != tc.want {
				t.Errorf("Sum(%v) = %d, want %d", tc.xs, got, tc.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		name   string
		xs     []int
		wantV  int
		wantOK bool
	}{
		{"nil", nil, 0, false},
		{"empty", []int{}, 0, false},
		{"single", []int{7}, 7, true},
		{"first_max", []int{9, 3, 1}, 9, true},
		{"last_max", []int{1, 3, 9}, 9, true},
		{"all_equal", []int{4, 4, 4}, 4, true},
		{"negatives", []int{-3, -1, -2}, -1, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gotV, gotOK := Max(tc.xs)
			if gotV != tc.wantV || gotOK != tc.wantOK {
				t.Errorf("Max(%v) = (%d, %v), want (%d, %v)",
					tc.xs, gotV, gotOK, tc.wantV, tc.wantOK)
			}
		})
	}
}

func TestContains(t *testing.T) {
	cases := []struct {
		name string
		xs   []int
		v    int
		want bool
	}{
		{"nil", nil, 1, false},
		{"empty", []int{}, 1, false},
		{"present", []int{1, 2, 3}, 2, true},
		{"absent", []int{1, 2, 3}, 9, false},
		{"first", []int{5, 6, 7}, 5, true},
		{"last", []int{5, 6, 7}, 7, true},
		{"negative_present", []int{-1, 0, 1}, -1, true},
		{"negative_absent", []int{-1, 0, 1}, -2, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Contains(tc.xs, tc.v); got != tc.want {
				t.Errorf("Contains(%v, %d) = %v, want %v",
					tc.xs, tc.v, got, tc.want)
			}
		})
	}
}
