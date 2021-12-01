package reverse

import (
	"unicode/utf8"
)

// Reverse function to reverse a String
func Reverse(s string) string {
	reversed := make([]rune, utf8.RuneCountInString(s))
	i := len(reversed) - 1

	for _, c := range s {
		reversed[i] = c
		i--
	}
	return string(reversed)
}
