package minesweeper

import (
	"fmt"
)

func Annotate(mines []string) []string {
	nbBombs := func(r, c int) string {
		bombs := 0
		adjacents := [][]int{{r - 1, c - 1}, {r - 1, c}, {r - 1, c + 1}, {r, c - 1}, {r, c + 1}, {r + 1, c - 1}, {r + 1, c}, {r + 1, c + 1}}
		for _, a := range adjacents {
			if a[0] >= 0 && a[0] < len(mines) && a[1] >= 0 && a[1] < len(mines[0]) && string(mines[a[0]][a[1]]) == "*" {
				bombs++
			}
		}
		if bombs == 0 {
			return " "
		}
		return fmt.Sprintf("%d", bombs)
	}

	annotated := make([]string, len(mines))
	for r, row := range mines {
		newRow := ""
		for c, char := range row {
			if string(char) == "*" {
				newRow += "*"

			} else {
				newRow += nbBombs(r, c)
			}
		}
		annotated[r] = newRow
	}
	return annotated
}
