package strutil_test

import (
	"testing"

	"github.com/ibalashov-srb/orch-test-target/strutil"
)

func TestCountVowels(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"rhythm", 0},
		{"aeiou", 5},
		{"AeIoU", 5},
		{"Hello World", 3},
		{"héllo", 1},
	}
	for _, c := range cases {
		if got := strutil.CountVowels(c.in); got != c.want {
			t.Errorf("CountVowels(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
