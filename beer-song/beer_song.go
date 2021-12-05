package beer

import (
	"errors"
	"fmt"
)

//Verse returns beer verse
func Verse(verse int) (string, error) {
	if verse < 0 || verse > 99 {
		return "", errors.New("Invalid verse")
	}
	if verse == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}
	if verse == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	}
	if verse == 2 {
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	}

	return fmt.Sprintf("%v bottles of beer on the wall, %v bottles of beer.\nTake one down and pass it around, %v bottles of beer on the wall.\n", verse, verse, verse-1), nil

}

//Verses returns beer verses
func Verses(from, to int) (string, error) {
	if from <= to {
		return "", errors.New("Invalid verses")
	}
	verses := ""
	for i := from; i >= to; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		verses += verse + "\n"
	}
	return verses, nil
}

//Song the whole song
func Song() string {
	song, _ := Verses(99, 0)
	return song
}
