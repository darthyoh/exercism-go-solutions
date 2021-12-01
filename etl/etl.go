package etl

import "strings"

//Transform functino
func Transform(input map[int][]string) map[string]int {

	mapped := make(map[string]int)

	for key, arr := range input {
		for _, v := range arr {
			mapped[strings.ToLower(v)] = key
		}
	}
	return mapped
}
