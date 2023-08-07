package bottlesong

import (
	"fmt"
	"strings"
)

var intMap map[int]string = map[int]string{
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

func Recite(startBottles, takeDown int) []string {
	if takeDown == 1 {
		return verse(startBottles)
	}
	verses := append(verse(startBottles), "")
	verses = append(verses, Recite(startBottles-1, takeDown-1)...)
	return verses
}

func verse(verse int) (verses []string) {
	if verse >= 2 {
		verses = []string{
			fmt.Sprintf("%s green bottles hanging on the wall,", intMap[verse]),
			fmt.Sprintf("%s green bottles hanging on the wall,", intMap[verse]),
		}
	} else {
		verses = []string{
			fmt.Sprintf("One green bottle hanging on the wall,"),
			fmt.Sprintf("One green bottle hanging on the wall,"),
		}
	}

	verses = append(verses, "And if one green bottle should accidentally fall,")

	switch {
	case verse > 2:
		verses = append(verses, fmt.Sprintf("There'll be %s green bottles hanging on the wall.", strings.ToLower(intMap[verse-1])))
	case verse == 2:
		verses = append(verses, "There'll be one green bottle hanging on the wall.")
	default:
		verses = append(verses, "There'll be no green bottles hanging on the wall.")
	}

	return

}
