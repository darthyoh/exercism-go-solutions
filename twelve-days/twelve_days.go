package twelve

import (
	"fmt"
	"strings"
)

type VerseContent struct {
	text    string
	content string
}

var versesMap = map[int]VerseContent{
	1: VerseContent{
		"first", "a Partridge in a Pear Tree",
	},
	2: VerseContent{
		"second", "two Turtle Doves",
	},
	3: VerseContent{
		"third", "three French Hens",
	},
	4: VerseContent{
		"fourth", "four Calling Birds",
	},
	5: VerseContent{
		"fifth", "five Gold Rings",
	},
	6: VerseContent{
		"sixth", "six Geese-a-Laying",
	},
	7: VerseContent{
		"seventh", "seven Swans-a-Swimming",
	},
	8: VerseContent{
		"eighth", "eight Maids-a-Milking",
	},
	9: VerseContent{
		"ninth", "nine Ladies Dancing",
	},
	10: VerseContent{
		"tenth", "ten Lords-a-Leaping",
	},
	11: VerseContent{
		"eleventh", "eleven Pipers Piping",
	},
	12: VerseContent{
		"twelfth", "twelve Drummers Drumming",
	},
}

func Song() string {
	string := ""
	for i := 1; i < 13; i++ {
		string += Verse(i)
		string += "\n"
	}
	return strings.TrimRight(string, "\n")
}

func Verse(n int) string {
	str := fmt.Sprintf("On the %v day of Christmas my true love gave to me: ", versesMap[n].text)
	if n == 1 {
		return fmt.Sprintf("%v%v.", str, versesMap[1].content)
	}
	contents := make([]string, 0)

	for i := n; i > 1; i-- {
		contents = append(contents, versesMap[i].content)
	}
	prefixe := strings.Join(contents, ", ")
	return fmt.Sprintf("%v%v, and %v.", str, prefixe, versesMap[1].content)

}
