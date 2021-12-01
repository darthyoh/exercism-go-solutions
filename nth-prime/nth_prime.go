package prime

//Nth returns the nth prime
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}

	nth := 1

	for i := 0; i < n; i++ {
		for {
			nth++
			if isPrime(nth) {
				break
			}
		}
	}

	return nth, true
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
