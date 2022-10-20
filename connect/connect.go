package connect

import (
	"strings"
)

type Position struct {
	row, col int
}

type Board struct {
	lines [][]string
}

func newBoard(lines []string) *Board {
	board := &Board{}

	for _, line := range lines {
		board.lines = append(board.lines, strings.Split(strings.ReplaceAll(line, " ", ""), ""))
	}
	return board
}

func (b *Board) charAt(p *Position) string {
	return b.lines[p.row][p.col]
}

func (p *Position) equals(p2 *Position) bool {
	return p.row == p2.row && p.col == p2.col
}

func (p *Position) adjacents(board *Board) (adj []*Position) {
	if p.col >= 1 {
		adj = append(adj, &Position{p.row, p.col - 1})
	}
	if p.col <= len(board.lines[0])-2 {
		adj = append(adj, &Position{p.row, p.col + 1})
	}
	if p.row >= 1 {
		adj = append(adj, &Position{p.row - 1, p.col})
		if p.col <= len(board.lines[0])-2 {
			adj = append(adj, &Position{p.row - 1, p.col + 1})
		}
		if p.col <= len(board.lines[0])-3 {
			adj = append(adj, &Position{p.row - 1, p.col + 2})
		}
	}
	if p.row <= len(board.lines)-2 {
		adj = append(adj, &Position{p.row + 1, p.col})
		if p.col >= 1 {
			adj = append(adj, &Position{p.row + 1, p.col - 1})
		}
		if p.col >= 2 {
			adj = append(adj, &Position{p.row + 1, p.col - 2})
		}

	}

	return
}

func (p *Position) hasReach(player string, board *Board) bool {
	if player == "X" {
		return p.col == len(board.lines[0])-1
	}
	return p.row == len(board.lines)-1
}

func ResultOf(lines []string) (winner string, err error) {
	board := newBoard(lines)

	var positionWins func(string, *Position, []*Position) bool

	positionWins = func(player string, p *Position, alreadyTested []*Position) bool {
		alreadyTested = append(alreadyTested, p)
		if p.hasReach(player, board) {
			return true
		}
		newAdjacents := make([]*Position, 0)
		for _, adjacent := range p.adjacents(board) {
			found := false
			for _, already := range alreadyTested {
				if adjacent.equals(already) {
					found = true
					break
				}
			}
			if !found {
				newAdjacents = append(newAdjacents, adjacent)
			}
		}
		for _, adjacent := range newAdjacents {
			if board.charAt(adjacent) != player {
				continue
			}
			if positionWins(player, adjacent, alreadyTested) {
				return true
			}
		}
		return false
	}

	for r := range board.lines {
		position := &Position{r, 0}
		if board.charAt(position) != "X" {
			continue
		}
		if positionWins("X", &Position{r, 0}, []*Position{}) {
			return "X", nil
		}
	}

	for c := range board.lines[0] {
		position := &Position{0, c}
		if board.charAt(position) != "O" {
			continue
		}
		if positionWins("O", &Position{0, c}, []*Position{}) {
			return "O", nil
		}
	}
	return
}
