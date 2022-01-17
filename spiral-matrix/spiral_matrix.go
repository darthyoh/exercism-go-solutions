package spiralmatrix

//SpiralMatrix constructs a matrix of n rows and cols
func SpiralMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	r, c := 0, 0
	dir := 'E'

	for i := 1; i <= n*n; i++ {
		matrix[r][c] = i
		switch dir {
		case 'E':
			if c == n-1 || matrix[r][c+1] != 0 {
				dir = 'S'
				r++
			} else {
				c++
			}
		case 'S':
			if r == n-1 || matrix[r+1][c] != 0 {
				dir = 'W'
				c--
			} else {
				r++
			}
		case 'W':
			if c == 0 || matrix[r][c-1] != 0 {
				dir = 'N'
				r--
			} else {
				c--
			}
		case 'N':
			if r == 0 || matrix[r-1][c] != 0 {
				dir = 'E'
				c++
			} else {
				r--
			}
		}
	}

	return matrix
}
