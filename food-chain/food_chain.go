package foodchain

import "fmt"

type verse struct {
	animal     string
	story      string
	complement string
}

var table = map[int]verse{
	1: verse{
		"fly", "I don't know why she swallowed the fly. Perhaps she'll die.", "",
	},
	2: verse{
		"spider", "It wriggled and jiggled and tickled inside her.", " that wriggled and jiggled and tickled inside her",
	},
	3: verse{
		"bird", "How absurd to swallow a bird!", "",
	},
	4: verse{
		"cat", "Imagine that, to swallow a cat!", "",
	},
	5: verse{
		"dog", "What a hog, to swallow a dog!", "",
	},
	6: verse{
		"goat", "Just opened her throat and swallowed a goat!", "",
	},
	7: verse{
		"cow", "I don't know how she swallowed a cow!", "",
	},
	8: verse{
		"horse", "She's dead, of course!", "",
	},
}

//Song gives the whole song !
func Song() string {
	return Verses(1, 8)
}

//Verse gives a verse of the song
func Verse(n int) string {
	sing := fmt.Sprintf("I know an old lady who swallowed a %v.\n%v", table[n].animal, table[n].story)
	if n == 8 || n == 1 {
		return sing
	}

	for i := n - 1; i >= 1; i-- {

		sing = fmt.Sprintf("%v\nShe swallowed the %v to catch the %v%v.", sing, table[i+1].animal, table[i].animal, table[i].complement)
	}

	sing = fmt.Sprintf("%v\n%v", sing, table[1].story)

	return sing

}

//Verses gives song from start to end verses
func Verses(start, end int) string {
	sing := Verse(start)

	for i := start + 1; i <= end; i++ {
		sing = fmt.Sprintf("%v\n\n%v", sing, Verse(i))
	}

	return sing
}
