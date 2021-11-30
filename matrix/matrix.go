package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	cols   int
	values []int
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 || row >= len(m.Rows()) || col >= len(m.Cols()) {
		return false
	}

	m.values[row*m.cols+col] = val
	return true
}

func (m *Matrix) Rows() [][]int {
	rows := make([][]int, 0)
	for i := 0; i < len(m.values)/m.cols; i++ {
		row := make([]int, 0)
		for j := 0; j < m.cols; j++ {
			row = append(row, m.values[i*m.cols+j])
		}
		rows = append(rows, row)
	}
	return rows
}

func (m *Matrix) Cols() [][]int {
	cols := make([][]int, 0)
	for i := 0; i < m.cols; i++ {
		col := make([]int, 0)
		for j := 0; j < len(m.values)/m.cols; j++ {
			col = append(col, m.values[j*m.cols+i])
		}
		cols = append(cols, col)
	}
	return cols

}

func New(s string) (*Matrix, error) {
	m := &Matrix{}
	values := make([]int, 0)

	lines := strings.Split(s, "\n")

	for _, line := range lines {
		line = strings.TrimLeft(line, " ")
		row := strings.Split(line, " ")

		if m.cols == 0 {
			m.cols = len(row)
		} else {
			if len(row) != m.cols {
				return nil, fmt.Errorf("Error : differents rows length")
			}
		}
		ints := make([]int, 0)
		for _, v := range row {
			if value, err := strconv.Atoi(v); err == nil {
				ints = append(ints, value)
			} else {
				return nil, fmt.Errorf("Error : not a int")
			}
		}
		values = append(values, ints...)

	}
	m.values = values
	return m, nil
}
