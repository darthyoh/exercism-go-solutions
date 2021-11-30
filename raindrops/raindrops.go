package raindrops

import (
	"fmt"
)

// Convert a value to PlingPlangPlong
func Convert(value int) string {
	if value%3 != 0 && value%5 != 0 && value%7 != 0 {
		return fmt.Sprint(value)
	}

	s := ""
	if value%3 == 0 {
		s += "Pling"
	}
	if value%5 == 0 {
		s += "Plang"
	}
	if value%7 == 0 {
		s += "Plong"
	}

	return s
}
