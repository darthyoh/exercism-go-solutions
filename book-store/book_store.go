package bookstore

import (
	"fmt"
	"sort"
)

var PRICES = map[int]int{
	1: 800,
	2: 1520,
	3: 2160,
	4: 2560,
	5: 3000,
}

func Cost(books []int) int {
	//get all possible combinaisons for books
	combinaisons := getCombinaisons12345(len(books))

	minCost := 0
	//test if any combinaison is possible
	for i := len(combinaisons) - 1; i >= 0; i-- {

		if price, err := getPriceForCombinaison(books, combinaisons[i]); err == nil && (minCost == 0 || price < minCost) {
			minCost = price
		}
	}

	return minCost

}

// simple Combinaison to sort a []int (in modern Go version, the slices package will do this instead)
type Combinaison []int

func (a Combinaison) Len() int {
	return len(a)
}
func (a Combinaison) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Combinaison) Less(i, j int) bool {
	return a[i] > a[j]
}

func (a Combinaison) Price() int {
	sum := 0

	for _, i := range a {
		cost, ok := PRICES[i]
		if ok {
			sum += cost
		}
	}
	return sum
}

// BookMap represents the number of each book
type BookMap map[int]int

func BookMapFromBooks(books []int) BookMap {
	bookMap := map[int]int{}
	for _, book := range books {
		if bks, ok := bookMap[book]; !ok {
			bookMap[book] = 1
		} else {
			bookMap[book] = bks + 1
		}
	}
	return bookMap
}

// RemoveGroup removes n books. It removes one entry of the most populated key
func (b BookMap) RemoveGroup(n int) {
	//sort keys in values desc

	keys := make([]int, 0, len(b))
	for k := range b {
		keys = append(keys, k)
	}
	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(keys)-i-1; j++ {
			if b[keys[j]] < b[keys[j+1]] {
				keys[j], keys[j+1] = keys[j+1], keys[j]
			}
		}
	}

	//delete one entry of each n first keys
	for i := 0; i < n; i++ {
		v := b[keys[i]]
		if v-1 == 0 {
			delete(b, keys[i])
		} else {
			b[keys[i]] = v - 1
		}
	}

}

// getPriceForCombinaison return, if combinaison is applicable to the books, the price or an error if not possible
func getPriceForCombinaison(books []int, combinaison []int) (int, error) {

	//sort combinaison in desc
	sort.Sort(Combinaison(combinaison))
	//construction of a map of book
	bookMap := BookMapFromBooks(books)

	for _, nbBooks := range combinaison {
		if len(bookMap) < nbBooks {
			return 0, fmt.Errorf("not applicable")
		}
		bookMap.RemoveGroup(nbBooks)
	}

	return Combinaison(combinaison).Price(), nil
}

// getCombinaisons12345 for a target returns all combinaison of differents books
func getCombinaisons12345(target int) [][]int {

	totalCombinaisons := make([][][]int, target+1)

	totalCombinaisons[0] = [][]int{{}}

	for intermediateTarget := 1; intermediateTarget < target+1; intermediateTarget++ {

		for nbBooks := 1; nbBooks < 6; nbBooks++ {

			if intermediateTarget-nbBooks >= 0 && totalCombinaisons[intermediateTarget-nbBooks] != nil {

				if totalCombinaisons[intermediateTarget] == nil {
					totalCombinaisons[intermediateTarget] = make([][]int, 0)
				}

				if len(totalCombinaisons[intermediateTarget-nbBooks]) == 0 {
					totalCombinaisons[intermediateTarget] = [][]int{{nbBooks}}
				} else {
					for _, intermediateTargetCombinaison := range totalCombinaisons[intermediateTarget-nbBooks] {
						newCombinaison := append([]int{}, intermediateTargetCombinaison...)
						newCombinaison = append(newCombinaison, nbBooks)

						totalCombinaisons[intermediateTarget] = append(totalCombinaisons[intermediateTarget], newCombinaison)
					}
				}

			}
		}
	}

	return totalCombinaisons[target]

}
