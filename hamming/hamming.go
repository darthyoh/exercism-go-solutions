package hamming

import "fmt"

//Distance func to get hamming distance between two DNA strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("Error")
	}
	differences := 0
	for i := range a {
		if a[i] != b[i] {
			differences++
		}
	}
	return differences, nil
}
