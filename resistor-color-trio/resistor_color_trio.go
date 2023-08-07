package resistorcolortrio

import (
	"fmt"
	"math"
	"regexp"
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

var kiloRe *regexp.Regexp
var megaRe *regexp.Regexp
var gigaRe *regexp.Regexp

func init() {
	kiloRe = regexp.MustCompile(`(\w+)000$`)
	megaRe = regexp.MustCompile(`(\w+)000000$`)
	gigaRe = regexp.MustCompile(`(\w+)000000000$`)
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	value := ""

	for i, v := range colors {
		if i <= 1 {
			value = fmt.Sprintf("%s%d", value, colorsMap[v])
		}
		if i == 2 {
			intValue, _ := strconv.Atoi(value)
			value = fmt.Sprintf("%d", int(float64(intValue)*math.Pow10(colorsMap[v])))
		}

	}

	if group := gigaRe.FindStringSubmatch(value); len(group) == 2 {
		return fmt.Sprintf("%s gigaohms", group[1])
	}

	if group := megaRe.FindStringSubmatch(value); len(group) == 2 {
		return fmt.Sprintf("%s megaohms", group[1])
	}

	if group := kiloRe.FindStringSubmatch(value); len(group) == 2 {
		return fmt.Sprintf("%s kiloohms", group[1])
	}

	return fmt.Sprintf("%s ohms", value)
}
