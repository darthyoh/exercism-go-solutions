// Package proverb for the Proverb function
package proverb

import "fmt"

// Proverb function
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return rhyme
	}

	var tab = make([]string, len(rhyme))

	for indice := range rhyme {
		if indice+1 < len(rhyme) {
			tab[indice] = fmt.Sprintf("For want of a %v the %v was lost.", rhyme[indice], rhyme[indice+1])

		} else {
			tab[indice] = fmt.Sprintf("And all for the want of a %v.", rhyme[0])
		}
	}

	return tab
}
