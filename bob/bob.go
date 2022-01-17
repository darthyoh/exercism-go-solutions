// Package bob for Hey function
package bob

import (
	"regexp"
	"strings"
)

// Hey function
func Hey(remark string) string {
	silence := regexp.MustCompile(`\s`)
	if silence.ReplaceAllString(remark, "") == "" {
		return "Fine. Be that way!"
	}

	remark = strings.Trim(remark, " ")

	nonLetters := regexp.MustCompile(`[^a-zA-Z]`)
	isQuestion := string(remark[len(remark)-1]) == "?"
	isYell := strings.ToUpper(remark) == remark && nonLetters.ReplaceAllString(remark, "") != ""

	if !isQuestion && !isYell {
		return "Whatever."
	} else if !isQuestion {
		return "Whoa, chill out!"
	} else if isYell {
		return "Calm down, I know what I'm doing!"
	} else {
		return "Sure."
	}
}
