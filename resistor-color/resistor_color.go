package resistorcolor

var colors map[string]int = map[string]int{
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

// Colors should return the list of all colors.
func Colors() []string {
	cs := make([]string, 0)

	for k := range colors {
		cs = append(cs, k)
	}

	return cs

}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	return colors[color]
}
