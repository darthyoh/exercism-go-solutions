package prime

//Factors return prime factors
func Factors(i int64) []int64 {
	factors := make([]int64, 0)
	if i == 1 {
		return factors
	}
	var activeFactor int64 = 2

	for {
		if i == 1 {
			break
		}
		if i%activeFactor == 0 {
			factors = append(factors, activeFactor)
			i = i / activeFactor
			continue
		}
		activeFactor++
	}

	return factors
}
