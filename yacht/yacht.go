package yacht

var funcs = map[string]func([]int) int{
	"yacht":           yacht,
	"ones":            ones,
	"twos":            twos,
	"threes":          threes,
	"fours":           fours,
	"fives":           fives,
	"sixes":           sixes,
	"full house":      full,
	"four of a kind":  four,
	"little straight": little,
	"big straight":    big,
	"choice":          choice,
}

func mapDice(dice []int) (mapp map[int]int, sum int) {
	mapp = map[int]int{}
	sum = 0
	for _, d := range dice {
		sum += d
		if _, ok := mapp[d]; !ok {
			mapp[d] = 0
		}
		mapp[d]++
	}
	return
}

func choice(dice []int) int {
	_, sum := mapDice(dice)
	return sum
}

func little(dice []int) int {
	if mapp, sum := mapDice(dice); len(mapp) == 5 && sum == 15 {
		return 30
	}
	return 0

}

func big(dice []int) int {
	if mapp, sum := mapDice(dice); len(mapp) == 5 && sum == 20 {
		return 30
	}
	return 0
}

func yacht(dice []int) int {
	if m, _ := mapDice(dice); len(m) == 1 {
		return 50
	}
	return 0
}

func unities(dice []int, unity int) int {
	m, _ := mapDice(dice)
	if v, ok := m[unity]; ok {
		return v * unity
	}
	return 0
}

func ones(dice []int) int {
	return unities(dice, 1)
}

func twos(dice []int) int {
	return unities(dice, 2)
}

func threes(dice []int) int {
	return unities(dice, 3)
}

func fours(dice []int) int {
	return unities(dice, 4)
}

func fives(dice []int) int {
	return unities(dice, 5)
}

func sixes(dice []int) int {
	return unities(dice, 6)
}

func full(dice []int) int {
	mapp, sum := mapDice(dice)

	if len(mapp) != 2 {
		return 0
	}

	for _, v := range mapp {
		if v != 2 && v != 3 {
			return 0
		}
	}
	return sum
}

func four(dice []int) int {
	mapp, _ := mapDice(dice)

	for k, v := range mapp {
		if v >= 4 {
			return 4 * k
		}
	}
	return 0
}

//Score returns score for a dice for a category in yacht game
func Score(dice []int, category string) int {
	if f, ok := funcs[category]; ok {
		return f(dice)
	}
	return 0
}
