package atbash

import (
	"regexp"
	"strings"
	"unicode"
)

//Atbash convert string to atbash cipher
func Atbash(input string) string {

	input = strings.ToLower(input)
	re := regexp.MustCompile(`[^a-z1-9]`)

	input = re.ReplaceAllString(input, "")

	var str strings.Builder

	letters := 0

	for i, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			if unicode.IsLetter(r) {
				str.WriteRune(219 - r)
			} else {
				str.WriteRune(r)
			}
			letters++
			if letters == 5 && i != len(input)-1 {
				str.WriteString(" ")
				letters = 0
			}
		}
	}

	return str.String()
}
