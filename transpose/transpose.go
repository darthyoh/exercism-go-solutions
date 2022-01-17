package transpose

import (
	"strings"
)

//Transpose converts horizontal to vertical matrix
func Transpose(matrix []string) []string {

	if len(matrix) == 0 {
		return []string{}
	}

	runes := make([][]rune, 0)
 
	for i, line := range matrix {
		for j, r := range line {
			if len(runes) <= j {
				runes = append(runes, make([]rune, len(matrix)))
			}
			runes[j][i] = r
		}
	}

	transposed := make([]string, 0)

	for _, line := range runes {
		transposed = append(transposed, string(line))
	}

	for i := len(transposed) - 1; i >= 0; i-- {
		if strings.TrimRight(transposed[i], "\x00") == transposed[i] {
			break
		}
		transposed[i] = strings.TrimRight(transposed[i], "\x00")
	}

	for i, line := range transposed {
		transposed[i] = strings.ReplaceAll(line, "\x00", " ")
	}

	return transposed
}
