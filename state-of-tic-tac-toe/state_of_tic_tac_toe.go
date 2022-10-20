package stateoftictactoe

import "fmt"

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	var countX, countO int

	for _, line := range board {
		for _, c := range line {
			if string(c) == "X" {
				countX++
			}
			if string(c) == "O" {
				countO++
			}
		}
	}

	if countO > countX {
		return "", fmt.Errorf("invalid board: O started")
	}

	if countX >= countO+2 {
		return "", fmt.Errorf("invalid board: X went twice")
	}

	rows := []string{
		board[0],
		board[1],
		board[2],
		fmt.Sprintf("%s%s%s", string(board[0][0]), string(board[1][0]), string(board[2][0])),
		fmt.Sprintf("%s%s%s", string(board[0][1]), string(board[1][1]), string(board[2][1])),
		fmt.Sprintf("%s%s%s", string(board[0][2]), string(board[1][2]), string(board[2][2])),
		fmt.Sprintf("%s%s%s", string(board[0][0]), string(board[1][1]), string(board[2][2])),
		fmt.Sprintf("%s%s%s", string(board[0][2]), string(board[1][1]), string(board[2][0])),
	}

	var xWin, oWin bool

	for _, row := range rows {
		if row == "XXX" {
			xWin = true
		}
		if row == "OOO" {
			oWin = true
		}
	}

	if xWin && oWin {
		return "", fmt.Errorf("invalid board: players kept playing after a win")
	}

	if !xWin && !oWin && countX+countO < 9 {
		return Ongoing, nil
	}

	if !xWin && !oWin && countX+countO == 9 {
		return Draw, nil
	}

	return Win, nil

}
