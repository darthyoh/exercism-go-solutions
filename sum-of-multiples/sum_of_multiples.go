package summultiples

//SumMultiples function
func SumMultiples(limit int, divisors ...int) int {

	sum := 0

	for i := range make([]int, limit-1) {
		for _, d := range divisors {
			if d > 0 && (i+1)%d == 0 {
				sum += (i + 1)
				break
			}
		}
	}

	return sum
}
