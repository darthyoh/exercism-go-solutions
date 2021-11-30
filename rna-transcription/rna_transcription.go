package strand

import "strings"

var transcription = map[rune]rune{
	71: 67,
	67: 71,
	84: 65,
	65: 85,
}

//ToRNA transciption from DNA
func ToRNA(dna string) string {

	var rna strings.Builder
	for _, r := range dna {
		rna.WriteRune(transcription[r])
	}
	return rna.String()
}
