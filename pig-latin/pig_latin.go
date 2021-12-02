package piglatin

import (
	"regexp"
	"strings"
)

//Sentence convert a sentence to pig latin
func Sentence(s string) string {
	words := strings.Fields(s)

	pig := make([]string, len(words))
	for i, w := range words {
		pig[i] = word(w)
	}
	return strings.Join(pig, " ")
}

//word convert a word to pig latin
func word(s string) string {
	re := regexp.MustCompile(`^(.+)y(.*)`)

	if re.MatchString(s) && s[len(s)-1:] != "y" {
		groups := re.FindStringSubmatch(s)
		return "y" + groups[2] + groups[1] + "ay"
	}

	re = regexp.MustCompile(`^[aeiou]|xr|yt`)
	if re.MatchString(s) {
		return s + "ay"
	}

	re = regexp.MustCompile(`^([bcdfghjklmnpqrstvwxz]qu|sch|ch|qu|thr|th|[bcdfghjklmnpqrstvwxyz])(.+)$`)
	if re.MatchString(s) {
		groups := re.FindStringSubmatch(s)
		return groups[2] + groups[1] + "ay"
	}

	return ""
}
