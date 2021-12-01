package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

//Encode function with crypto square
func Encode(s string) string {

	if s == "" {
		return ""
	}
	var cleaned strings.Builder
	var encoded strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleaned.WriteRune(unicode.ToLower(r))
		}
	}

	s = cleaned.String()

	cols := int(math.Ceil(math.Sqrt(float64(len(s)))))
	rows := int(math.Ceil(float64(len(s)) / float64(cols)))

	table := make([][]rune, rows)

	line := make([]rune, cols)

	activeRow := 0

	for i, r := range s {
		line[i%cols] = r
		if (i+1)%cols == 0 || i == len(s)-1 {
			table[activeRow] = line
			activeRow++
			line = make([]rune, cols)
		}
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if table[j][i] == 0 {
				encoded.WriteString(" ")
			} else {
				encoded.WriteRune(table[j][i])
			}
		}

		if i < cols-1 {
			encoded.WriteString(" ")
		}

	}
	return encoded.String()
}
