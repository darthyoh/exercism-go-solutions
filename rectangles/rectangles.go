package rectangles

type rectangle struct {
	x, y int
}

func (r *rectangle) isValid(newX, newY int, input *[]string) bool {
	if (*input)[r.y][newX] != 43 || (*input)[newY][r.x] != 43 {
		return false
	}

	for x := r.x + 1; x < newX; x++ {
		if ((*input)[r.y][x] != 43 && (*input)[r.y][x] != 45) || ((*input)[newY][x] != 43 && (*input)[newY][x] != 45) {
			return false
		}
	}

	for y := r.y + 1; y < newY; y++ {
		if ((*input)[y][r.x] != 43 && (*input)[y][r.x] != 124) || ((*input)[y][newX] != 43 && (*input)[y][newX] != 124) {
			return false
		}
	}

	return true
}

//Count returns the number of rectangle in an input string
func Count(input []string) int {
	counter := 0
	for y, line := range input {
		for x, r := range line {
			if r == 43 {
				newRect := rectangle{x: x, y: y}
				if y < len(input)-1 && x < len(line)-1 {
					for newX := x + 1; newX < len(line); newX++ {
						for newY := y + 1; newY < len(input); newY++ {
							if input[newY][newX] == 43 && newRect.isValid(newX, newY, &input) {
								counter++
							}
						}
					}
				}
			}
		}
	}
	return counter
}
