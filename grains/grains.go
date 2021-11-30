package grains

import (
	"fmt"
)

//Square calculates the grains of a case
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, fmt.Errorf("Error")
	}

	return uint64(1 << (n - 1)), nil
}

func total(n int) uint64 {
	if n == 1 {
		return 1
	}
	tot, _ := Square(n)
	return tot + total(n-1)
}

//Total calculates the grains on a Chessboard
func Total() uint64 {
	return total(64)
}
