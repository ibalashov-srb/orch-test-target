package sliceutil_test

import (
	"testing"

	"github.com/ibalashov-srb/orch-test-target/sliceutil"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		want int
	}{
		{name: "nil slice", xs: nil, want: 0},
		{name: "empty slice", xs: []int{}, want: 0},
		{name: "mixed signs sum to zero", xs: []int{-1, -2, 3}, want: 0},
		{name: "all same positive", xs: []int{2, 2, 2}, want: 6},
		{name: "single element", xs: []int{42}, want: 42},
		{name: "all negative", xs: []int{-3, -7}, want: -10},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sliceutil.Sum(tc.xs)
			if got != tc.want {
				t.Errorf("Sum(%v) = %d; want %d", tc.xs, got, tc.want)
			}
		})
	}
}
