package anagram

import (
	"sort"
	"strings"
)

type sortedString struct {
	initial string
	sorted  string
}

func newSortedString(input string) sortedString {
	upper := strings.ToUpper(input)
	sorted := strings.Split(upper, "")
	sort.Strings(sorted)
	return sortedString{input, strings.Join(sorted, "")}
}

//Detect anagrams for input from candidates
func Detect(input string, candidates []string) (anagrams []string) {

	sortedInput := newSortedString(input)
	anagrams = make([]string, 0)

	for _, candidate := range candidates {
		sortedCandidate := newSortedString(candidate)
		if strings.ToUpper(sortedCandidate.initial) != strings.ToUpper(input) {
			if sortedCandidate.sorted == sortedInput.sorted {
				anagrams = append(anagrams, candidate)
			}
		}
	}
	return
}
