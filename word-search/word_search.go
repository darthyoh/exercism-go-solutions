package wordsearch

import (
	"errors"
)

type dir int

//Const definitions
const (
	E dir = iota
	W
	N
	S
	SE
	NE
	SW
	NW
)

var directions = []dir{E, W, N, S, SE, NE, SW, NW}

//Puzzle type
type Puzzle []string

func (p *Puzzle) getWordLine(col, row int, dir dir) string {
	sizeX := len((*p)[0])
	w := ""
	switch dir {
	case E:
		for i := col; i < sizeX; i++ {
			w += string((*p)[row][i])
		}

	case W:

		for i := col; i >= 0; i-- {
			w += string((*p)[row][i])
		}

	case S:

		for i := row; i < len(*p); i++ {
			w += string((*p)[i][col])
		}

	case N:

		for i := row; i > 0; i-- {
			w += string((*p)[i][col])
		}

	case NE:

		for i, j := row, col; i > 0 && j < sizeX; i, j = i-1, j+1 {
			w += string((*p)[i][j])
		}

	case NW:

		for i, j := row, col; i > 0 && j > 0; i, j = i-1, j-1 {
			w += string((*p)[i][j])
		}

	case SE:

		for i, j := row, col; i < len((*p)) && j < sizeX; i, j = i+1, j+1 {
			w += string((*p)[i][j])
		}

	case SW:

		for i, j := row, col; i < len((*p)) && j > 0; i, j = i+1, j-1 {
			w += string((*p)[i][j])
		}

	}

	return w
}

func (p *Puzzle) getPosition(word string) (positions [2][2]int, error error) {

	error = errors.New("Not found")
	positions = [2][2]int{}
	sizeX := len((*p)[0])

	for i := 0; i < len(*p); i++ {
		for j := 0; j < sizeX; j++ {
			for _, dir := range directions {
				w := p.getWordLine(j, i, dir)
				if len(w) >= len(word) && word == w[0:len(word)] {
					switch dir {
					case N:
						positions, error = [2][2]int{{j, i}, {j, i - (len(word) - 1)}}, nil
					case S:
						positions, error = [2][2]int{{j, i}, {j, i + (len(word) - 1)}}, nil
					case E:
						positions, error = [2][2]int{{j, i}, {j + (len(word) - 1), i}}, nil
					case W:
						positions, error = [2][2]int{{j, i}, {j - (len(word) - 1), i}}, nil
					case NE:
						positions, error = [2][2]int{{j, i}, {j + (len(word) - 1), i - (len(word) - 1)}}, nil
					case NW:
						positions, error = [2][2]int{{j, i}, {j - (len(word) - 1), i - (len(word) - 1)}}, nil
					case SE:
						positions, error = [2][2]int{{j, i}, {j + (len(word) - 1), i + (len(word) - 1)}}, nil
					case SW:
						positions, error = [2][2]int{{j, i}, {j - (len(word) - 1), i + (len(word) - 1)}}, nil
					}
					return
				}
			}
		}
	}
	return
}

//Solve returns a map with positions found for each word in a puzzle
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {

	solutions := map[string][2][2]int{}

	var p Puzzle = puzzle

	for _, word := range words {
		positions, err := p.getPosition(word)
		if err != nil {
			return solutions, errors.New("word not found")
		}
		solutions[word] = positions

	}

	if len(solutions) != len(words) {
		return solutions, errors.New("No complete solution")
	}

	return solutions, nil
}
