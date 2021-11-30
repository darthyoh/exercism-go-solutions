package luhn

import (
	"fmt"
	"regexp"
	"unicode"
)
//Valid tests if a code is a valid Luhn number
func Valid(code string) bool {
	re := regexp.MustCompile(`\s`)
	code = re.ReplaceAllString(code, "")
	if len(code) <= 1 {
		return false
	}

	sum := 0

	reversed := ""

	for _, runeValue := range code {
		reversed = fmt.Sprintf("%c", runeValue) + reversed
	}

	for i, runeValue := range reversed {
		if !unicode.IsNumber(runeValue) {
			return false
		}
		digit := int(runeValue - '0')

		if i%2 != 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}
	return sum%10 == 0
}
