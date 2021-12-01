package isbn

import (
	"regexp"
	"strings"
)

//IsValidISBN checks if string is a valid ISBN number
func IsValidISBN(s string) bool {

	if matched, err := regexp.MatchString(`(^[0-9]{1}-[0-9]{3}-[0-9]{5}-([0-9]{1}|X)$)|(^[0-9]{9}([0-9]{1}|X)$)`, s); !matched || err != nil {
		return false
	}

	s = strings.ReplaceAll(s, "-", "")

	sum := 0

	for i, r := range s {

		value := int(r - '0')
		if i == 9 && r == 88 {
			value = 10
		}

		sum += (10 - i) * value

	}

	return sum%11 == 0
}
