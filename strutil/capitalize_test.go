package strutil_test

import (
	"testing"

	"github.com/ibalashov-srb/orch-test-target/strutil"
)

func TestCapitalize(t *testing.T) {
	cases := []struct{ in, want string }{
		{"", ""},
		{"hello", "Hello"},
		{"Hello", "Hello"},
		{"hello world", "Hello World"},
		{"hELLO wORLD", "HELLO WORLD"},
		{" hello", " Hello"},
		{"héllo", "Héllo"},
		{"hello  world", "Hello  World"}, // double space: extra empty token preserved
	}
	for _, c := range cases {
		if got := strutil.Capitalize(c.in); got != c.want {
			t.Errorf("Capitalize(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
