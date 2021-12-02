package railfence

import "fmt"

//Encode string with rail cipher
func Encode(input string, rail int) (encoded string) {
	tables := make([][]rune, rail)
	for i := range tables {
		tables[i] = make([]rune, len(input))
	}

	row := 0
	direction := 1

	for i, r := range input {
		tables[row][i] = r
		switch direction {
		case 1:
			if row < len(tables)-1 {
				row++
			} else {
				direction = -1
				row--
			}
		case -1:
			if row > 0 {
				row--
			} else {
				direction = 1
				row++
			}
		}
	}

	encoded = ""

	for _, table := range tables {
		for _, r := range table {
			if r != 0 {
				encoded = fmt.Sprintf("%v%c", encoded, r)
			}
		}
	}

	return
}

//Decode string with rail cipher
func Decode(input string, rail int) (decoded string) {
	tables := make([][]rune, rail)
	for i := range tables {
		tables[i] = make([]rune, len(input))
	}

	row := 0
	direction := 1

	for i := range input {
		tables[row][i] = 1
		switch direction {
		case 1:
			if row < len(tables)-1 {
				row++
			} else {
				direction = -1
				row--
			}
		case -1:
			if row > 0 {
				row--
			} else {
				direction = 1
				row++
			}
		}
	}

	for i, table := range tables {
		for j, r := range table {
			if r == 1 {
				tables[i][j] = rune(input[0])
				input = input[1:]
			}
		}
	}

	decoded = ""

	row = 0
	direction = 1

	for i := range tables[0] {
		decoded = fmt.Sprintf("%v%c", decoded, tables[row][i])
		switch direction {
		case 1:
			if row < len(tables)-1 {
				row++
			} else {
				direction = -1
				row--
			}
		case -1:
			if row > 0 {
				row--
			} else {
				direction = 1
				row++
			}
		}
	}

	return
}
