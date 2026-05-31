package strutil

import (
	"strings"
	"unicode"
)

// Capitalize returns s with the first rune of each space-separated word
// uppercased, leaving the remainder of each word unchanged.
func Capitalize(s string) string {
	tokens := strings.Split(s, " ")
	for i, token := range tokens {
		if token == "" {
			continue
		}
		runes := []rune(token)
		runes[0] = unicode.ToUpper(runes[0])
		tokens[i] = string(runes)
	}
	return strings.Join(tokens, " ")
}
