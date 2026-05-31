package strutil

import (
	"strings"
	"unicode"
)

// IsPalindrome reports whether s is a palindrome, ignoring spaces and case.
func IsPalindrome(s string) bool {
	normalized := strings.Map(unicode.ToLower, strings.ReplaceAll(s, " ", ""))
	return normalized == Reverse(normalized)
}
