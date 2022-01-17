package scale

import "strings"

var sharpChromatic = [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var flatChromatic = [12]string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}

func getChromaticForTonic(tonic string) []string {

	var chromatic *[12]string

	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#", "a":
		chromatic = &sharpChromatic
	default:
		chromatic = &flatChromatic
	}

	start := 0

	for i := 0; i < 12; i++ {
		if chromatic[i] == strings.ToUpper(tonic[0:1])+tonic[1:] {
			start = i
			break
		}
	}

	scale := make([]string, 12)

	for i := 0; i < 12; i++ {
		scale[i] = chromatic[start]
		if start == 11 {
			start = 0
		} else {
			start++
		}
	}

	return scale
}

//Scale generator
func Scale(tonic string, interval string) []string {

	chromatic := getChromaticForTonic(tonic)

	if len(interval) == 0 { // chromatic
		return chromatic
	}

	scale := make([]string, 0)

	scale = append(scale, strings.ToUpper(tonic[0:1])+tonic[1:])
	offset := 0

	for _, r := range interval {
		switch string(r) {
		case "M":
			offset += 2
		case "m":
			offset++
		case "A":
			offset += 3
		}

		scale = append(scale, chromatic[offset%12])
	}

	return scale[:len(scale)-1]
}
