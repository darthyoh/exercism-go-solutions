package house

import "fmt"

type element struct {
	name string
	do   string
}

var elements = map[int]element{
	1:  element{"the house that Jack built", ""},
	2:  element{"the malt", "lay in"},
	3:  element{"the rat", "ate"},
	4:  element{"the cat", "killed"},
	5:  element{"the dog", "worried"},
	6:  element{"the cow with the crumpled horn", "tossed"},
	7:  element{"the maiden all forlorn", "milked"},
	8:  element{"the man all tattered and torn", "kissed"},
	9:  element{"the priest all shaven and shorn", "married"},
	10: element{"the rooster that crowed in the morn", "woke"},
	11: element{"the farmer sowing his corn", "kept"},
	12: element{"the horse and the hound and the horn", "belonged to"},
}

//Verse gives a verse of the Song
func Verse(verse int) string {
	if verse == 1 {
		return fmt.Sprintf("This is %v.", elements[1].name)
	}

	s := fmt.Sprintf("This is %v", elements[verse].name)
	for i := verse; i > 1; i-- {
		s = fmt.Sprintf("%v\nthat %v %v", s, elements[i].do, elements[i-1].name)
	}
	return s + "."

}

//Song gives the whole song
func Song() string {
	song := Verse(1)
	for i := 2; i < 13; i++ {
		song = fmt.Sprintf("%v\n\n%v", song, Verse(i))
	}
	return song
}
