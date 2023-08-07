package resistorcolorduo

import (
	"fmt"
	"strconv"
)

var colorsMap map[string]int = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	value := ""

	for i, v := range colors {
		if i > 1 {
			break
		}
		value = fmt.Sprintf("%s%d", value, colorsMap[v])
	}

	if intValue, err := strconv.Atoi(value); err != nil {
		return 0
	} else {
		return intValue
	}
}
