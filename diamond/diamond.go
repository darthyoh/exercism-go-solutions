package diamond

import (
	"errors"
	"fmt"
	"strings"
)

//Gen diamond from letter
func Gen(input byte) (string, error) {
	if input < 65 || input > 90 {
		return "", errors.New("Invalid byte")
	}

	//init strings array
	size := (int(input)-65)*2 + 1
	lines := make([]string, size)

	//fill first and last lines
	lines[0] = fmt.Sprintf("%vA%v", strings.Repeat(" ", size/2), strings.Repeat(" ", size/2))
	lines[size-1] = lines[0]

	//fill middle lines
	for i := 1; i <= (size / 2); i++ {
		lines[i] = fmt.Sprintf("%v%c%v%c%v", strings.Repeat(" ", (size/2)-i), 65+i, strings.Repeat(" ", i*2-1), 65+i, strings.Repeat(" ", (size/2)-i))
		lines[size-1-i] = lines[i]
	}

	return strings.Join(lines, "\n") + "\n", nil
}
