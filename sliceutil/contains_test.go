package sliceutil_test

import (
	"testing"

	"github.com/ibalashov-srb/orch-test-target/sliceutil"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		v    int
		want bool
	}{
		{name: "nil slice yields false", xs: nil, v: 1, want: false},
		{name: "empty slice yields false", xs: []int{}, v: 0, want: false},
		{name: "element found at start", xs: []int{-1, 0, 1}, v: -1, want: true},
		{name: "element found in middle", xs: []int{-1, 0, 1}, v: 0, want: true},
		{name: "element found at end", xs: []int{-1, 0, 1}, v: 1, want: true},
		{name: "element not found", xs: []int{1, 2, 3}, v: 9, want: false},
		{name: "single element match", xs: []int{42}, v: 42, want: true},
		{name: "single element no match", xs: []int{42}, v: 0, want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := sliceutil.Contains(tc.xs, tc.v)
			if got != tc.want {
				t.Errorf("Contains(%v, %d) = %v; want %v", tc.xs, tc.v, got, tc.want)
			}
		})
	}
}
