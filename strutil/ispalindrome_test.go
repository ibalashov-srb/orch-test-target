package strutil_test

import (
	"testing"

	"github.com/ibalashov-srb/orch-test-target/strutil"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{"a", true},
		{"RaceCar", true},
		{"A man a plan a canal Panama", true},
		{"hello", false},
	}
	for _, c := range cases {
		if got := strutil.IsPalindrome(c.in); got != c.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
