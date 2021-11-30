package isogram

import (
	"unicode"
)

// IsIsogram determine if the string is an isogram...
func IsIsogram(s string) bool {
	table := make(map[rune]int)
	for _, runeValue := range s {
		if unicode.IsLetter(runeValue) {
			runeValue = unicode.ToUpper(runeValue)
			if _, present := table[runeValue]; !present {
				table[runeValue] = 0
			}
			table[runeValue]++
		}

	}

	for _, value := range table {
		if value != 1 {
			return false
		}
	}

	return true
}
