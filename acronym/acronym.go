// Package acronym ...
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate function
func Abbreviate(s string) string {
	notAChar := regexp.MustCompile(`[^a-zA-Z']`)
	s = notAChar.ReplaceAllString(s, " ")
	acro := ""
	for _, word := range strings.Fields(s) {
		acro += word[:1]
	}

	return strings.ToUpper(acro)
}
