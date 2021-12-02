package pascal

//Triangle function
func Triangle(line int) [][]int {
	if line == 1 {
		return [][]int{{1}}
	}

	triangle := Triangle(line - 1)

	lastLine := triangle[line-2]
	newLine := make([]int, line)
	newLine[0], newLine[line-1] = 1, 1

	for i := 1; i <= line-2; i++ {
		newLine[i] = lastLine[i] + lastLine[i-1]
	}

	return append(triangle, newLine)
}
