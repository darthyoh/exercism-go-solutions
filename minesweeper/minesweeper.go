package minesweeper

import (
	"errors"
	"regexp"
	"strings"
)

type position struct{ row, col uint }

//Count generates the solution
func (b *Board) Count() error {
	initialRe := regexp.MustCompile(`^\+[-]*\+$`)
	centralRe := regexp.MustCompile(`^\|[ \*]*\|$`)
	rows := strings.Split(b.String(), "\n")[1:]

	var size = 0
	for i, row := range rows {
		if i == 0 {
			size = len(row)
		} else {
			if len(row) != size {
				return errors.New("Invalid board cause all rows are not equal")
			}
		}
		if i == 0 || i == len(rows)-1 {
			if !initialRe.MatchString(row) {
				return errors.New("Invalid board cause first or last row is not valid")
			}
		} else {
			if !centralRe.MatchString(row) {
				return errors.New("Invalid board cause one line is not valid")
			}
		}
	}

	for r := 1; r < len(*b)-1; r++ {
		for c := 1; c < size-1; c++ {
			if (*b)[r][c] != byte('*') {
				var bombs uint8 = 0
				for i := c - 1; i <= c+1; i++ {
					if (*b)[r-1][i] == byte('*') {
						bombs++
					}
					if (*b)[r+1][i] == byte('*') {
						bombs++
					}
					if (*b)[r][i] == byte('*') && i != c {
						bombs++
					}
				}
				if bombs != 0 {
					(*b)[r][c] = byte(48 + bombs)
				}
			}
		}
	}

	return nil
}
