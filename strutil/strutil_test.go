package strutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: ""},
		{name: "single", input: "a", want: "a"},
		{name: "ascii", input: "hello", want: "olleh"},
		{name: "unicode two-byte", input: "héllo", want: "olléh"},
		{name: "unicode multi-rune", input: "日本語", want: "語本日"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Reverse(tc.input); got != tc.want {
				t.Errorf("Reverse(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "empty", input: "", want: true},
		{name: "single", input: "a", want: true},
		{name: "simple true", input: "racecar", want: true},
		{name: "simple false", input: "hello", want: false},
		{name: "mixed case", input: "RaceCar", want: true},
		{name: "with spaces", input: "A man a plan a canal Panama", want: true},
		{name: "non-palindrome with spaces", input: "not a palindrome", want: false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsPalindrome(tc.input); got != tc.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestCountVowels(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{name: "empty", input: "", want: 0},
		{name: "no vowels", input: "rhythm", want: 0},
		{name: "all vowels", input: "aeiou", want: 5},
		{name: "mixed case", input: "AeIoU", want: 5},
		{name: "sentence", input: "Hello World", want: 3},
		{name: "unicode non-vowel", input: "héllo", want: 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountVowels(tc.input); got != tc.want {
				t.Errorf("CountVowels(%q) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "empty", input: "", want: ""},
		{name: "single word", input: "hello", want: "Hello"},
		{name: "already capitalized", input: "Hello", want: "Hello"},
		{name: "multi-word", input: "hello world", want: "Hello World"},
		{name: "mixed case", input: "hELLO wORLD", want: "HELLO WORLD"},
		{name: "leading space", input: " hello", want: " Hello"},
		{name: "unicode first rune", input: "héllo", want: "Héllo"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Capitalize(tc.input); got != tc.want {
				t.Errorf("Capitalize(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
