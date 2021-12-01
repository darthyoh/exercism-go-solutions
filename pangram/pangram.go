package pangram

import (
	"unicode"
)

//IsPangram determines if a string (s) is a Pangram
func IsPangram(s string) bool {

	table := make(map[rune]int)

	for _, v := range s {
		if unicode.IsLetter(v) {
			v = unicode.ToLower(v)
			if _, ok := table[v]; !ok {
				table[v] = 0
			}
			table[v]++
		}
	}

	return len(table) == 26
}
