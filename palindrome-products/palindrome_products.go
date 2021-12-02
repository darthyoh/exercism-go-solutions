package palindrome

import (
	"errors"
	"sort"
	"strconv"
)

//Product struct for storing factorizations of a palindromic product
type Product struct {
	Product        int
	Factorizations [][2]int
}

func (p *Product) addFacto(i, j int) {

	if i > j {
		i, j = j, i
	}

	for _, v := range p.Factorizations {
		if v == [2]int{i, j} {
			return
		}
	}

	p.Factorizations = append(p.Factorizations, [2]int{i, j})
}

//Products generates min and max palindromic product
func Products(fmin, fmax int) (Product, Product, error) {

	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax...")
	}

	var products = make(map[int]*Product)

	for i := fmin; i <= fmax; i++ {
		for j := fmin; j <= fmax; j++ {
			if isPalindromic(i * j) {
				if _, ok := products[i*j]; !ok {
					products[i*j] = &Product{i * j, make([][2]int, 0)}
				}

				products[i*j].addFacto(i, j)
			}
		}
	}

	keys := make([]int, 0)

	for k := range products {
		keys = append(keys, k)
	}

	if len(keys) == 0 {
		return Product{}, Product{}, errors.New("no palindromes...")
	}

	sort.Ints(keys)

	return *products[keys[0]], *products[keys[len(keys)-1]], nil

}

func isPalindromic(i int) bool {
	if i < 10 {
		return true
	}
	s := strconv.Itoa(i)
	reversed := ""
	for _, v := range s {
		reversed = string(v) + reversed
	}

	if v, err := strconv.Atoi(reversed); err == nil {
		return v == i
	}

	return false
}
