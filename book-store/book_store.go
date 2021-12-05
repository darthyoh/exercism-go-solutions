package bookstore

import (
	"sync"
)

var groups map[int]int

func init() {

	groups = map[int]int{
		5: 5 * 800 * 0.75,
		4: 4 * 800 * 0.80,
		3: 3 * 800 * 0.90,
		2: 2 * 800 * 0.95,
		1: 800,
	}
}

func Cost(books []int) int {
	min := 0

	sums := getAllSolutions(books)

	for v := range sums {
		if min == 0 || v < min {
			min = v
		}
	}
	return min
}

func getAllSolutions(books []int) chan int {
	sums := make(chan int)
	go func() {
		defer close(sums)
		//get a map of books
		differentBooks := make(map[int]int)

		for _, b := range books {
			if _, ok := differentBooks[b]; !ok {
				differentBooks[b] = 0
			}
			differentBooks[b]++
		}

		if len(differentBooks) < 2 {
			//Case no book, or 1 only book or only same books
			sums <- len(books) * 800
		} else {
			//search for max grouping possibility (limited to 5)
			maxGroup := len(differentBooks)
			if maxGroup > 5 {
				maxGroup = 5
			}

			for group := maxGroup; group > 0; group-- {
				//get keys of Map
				keys := make([]int, 0)
				for k := range differentBooks {
					keys = append(keys, k)
				}

				var wg sync.WaitGroup

				//get all solutions for these keys
				for _, possibility := range getPossiblesGroups(keys, group) {

					wg.Add(1)

					possibility := possibility

					func() {
						defer wg.Done()
						//getting left books
						restingBooks := make([]int, 0)
						for i, v := range keys {
							bookToRemove := 0
							if possibility[i] {
								bookToRemove = 1
							}
							for j := 0; j < differentBooks[v]-bookToRemove; j++ {
								restingBooks = append(restingBooks, v)
							}
						}
						for newSum := range getAllSolutions(restingBooks) {
							sums <- groups[group] + newSum
						}
					}()
				}
				wg.Wait()
			}
		}
	}()
	return sums

}

func getPossiblesGroups(input []int, group int) [][]bool {
	if len(input) == 0 || group > len(input) {
		return [][]bool{}
	}
	if len(input) == group {
		solution := make([]bool, 0)
		for range input {
			solution = append(solution, true)
		}
		return [][]bool{solution}
	}

	solutions := make([][]bool, 0)

	if group == 1 {
		for i := range input {
			solution := make([]bool, len(input))
			solution[i] = true
			solutions = append(solutions, solution)
		}
		return solutions
	}
	for i := 0; i < len(input)-group+1; i++ {
		solution := make([]bool, i+1)
		solution[i] = true
		subsolutions := getPossiblesGroups(input[i+1:], group-1)
		for _, v := range subsolutions {
			solutions = append(solutions, append(solution, v...))
		}
	}
	return solutions
}
