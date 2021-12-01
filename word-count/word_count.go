package wordcount

import (
	"regexp"
	"strings"
)

//Frequency is a map of counting words
type Frequency map[string]int

//WordCount gives a Frequency count of each word
func WordCount(input string) Frequency {
	input = strings.ToLower(input)

	invalidCar := regexp.MustCompile(`[^a-zA-Z0-9']`)
	quote := regexp.MustCompile(`^'|'$`)

	input = invalidCar.ReplaceAllString(input, " ")
	words := strings.Fields(input)

	frequency := make(Frequency)

	for _, word := range words {
		word = quote.ReplaceAllString(word, "")
		if _, ok := frequency[word]; !ok {
			frequency[word] = 0
		}
		frequency[word]++
	}

	return frequency
}
