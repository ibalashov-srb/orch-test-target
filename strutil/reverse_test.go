package strutil_test

import (
	"testing"

	"github.com/ibalashov-srb/orch-test-target/strutil"
)

func TestReverse(t *testing.T) {
	cases := []struct{ in, want string }{
		{"", ""},
		{"hello", "olleh"},
		{"日本語", "語本日"},
	}
	for _, c := range cases {
		if got := strutil.Reverse(c.in); got != c.want {
			t.Errorf("Reverse(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
