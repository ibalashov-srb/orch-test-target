package sliceutil_test

import (
	"testing"

	"github.com/ibalashov-srb/orch-test-target/sliceutil"
)

func TestMax(t *testing.T) {
	tests := []struct {
		name    string
		xs      []int
		wantVal int
		wantOK  bool
	}{
		{name: "nil slice", xs: nil, wantVal: 0, wantOK: false},
		{name: "empty slice", xs: []int{}, wantVal: 0, wantOK: false},
		{name: "single element", xs: []int{7}, wantVal: 7, wantOK: true},
		{name: "all negative", xs: []int{-3, -1, -2}, wantVal: -1, wantOK: true},
		{name: "mixed signs", xs: []int{-5, 0, 3, 1}, wantVal: 3, wantOK: true},
		{name: "all same", xs: []int{4, 4, 4}, wantVal: 4, wantOK: true},
		{name: "ascending order", xs: []int{1, 2, 3, 4, 5}, wantVal: 5, wantOK: true},
		{name: "descending order", xs: []int{5, 4, 3, 2, 1}, wantVal: 5, wantOK: true},
		{name: "max at start", xs: []int{10, 1, 2}, wantVal: 10, wantOK: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotVal, gotOK := sliceutil.Max(tc.xs)
			if gotVal != tc.wantVal || gotOK != tc.wantOK {
				t.Errorf("Max(%v) = (%d, %v); want (%d, %v)",
					tc.xs, gotVal, gotOK, tc.wantVal, tc.wantOK)
			}
		})
	}
}
