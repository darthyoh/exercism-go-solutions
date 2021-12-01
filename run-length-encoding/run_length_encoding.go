package encode

import "fmt"
import "unicode"
import "strings"
import "strconv"

//RunLengthEncode decodes string
func RunLengthEncode(s string) string {
	if s == "" {
		return ""
	}

	encoded := ""
	lastRune := rune(0)
	count := 0

	encode := func() {
		if count != 1 {
			encoded += fmt.Sprintf("%v", count)
		}
		encoded += fmt.Sprintf("%c", lastRune)
	}

	for _, r := range s {
		if lastRune == 0 {
			count = 1
			lastRune = r
			continue
		}
		if r == lastRune {
			count++
		} else {
			encode()
			count = 1
			lastRune = r
		}
	}
	encode()

	return encoded
}

//RunLengthDecode decodes string
func RunLengthDecode(s string) string {
	if s == "" {
		return ""
	}

	decoded := ""
	count := 0

	for _, r := range s {

		if unicode.IsDigit(r) {
			value, _ := strconv.Atoi(string(r))
			if count == 0 {
				count = value
			} else {
				newCount, _ := strconv.Atoi(fmt.Sprintf("%v%v", count, value))
				count = newCount
			}
			continue
		}
		if count == 0 {
			decoded += fmt.Sprintf("%c", r)
			continue
		}
		decoded += strings.Repeat(string(r), count)
		count = 0
	}
	return decoded
}
