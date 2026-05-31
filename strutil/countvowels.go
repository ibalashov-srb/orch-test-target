package strutil

import "unicode"

// CountVowels returns the number of ASCII vowels (a, e, i, o, u) in s.
// Case-insensitive; non-ASCII runes are not counted.
func CountVowels(s string) int {
	count := 0
	for _, r := range s {
		switch unicode.ToLower(r) {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
