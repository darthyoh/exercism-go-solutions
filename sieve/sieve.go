package sieve

//Sieve returns all primes
func Sieve(n int) []int {
	if n < 2 {
		return []int{}
	}
	table := make([]bool, n+1)

	primes := make([]int, 0)

	for i := 2; i <= n; i++ {
		if table[i] == false {
			primes = append(primes, i)
			for j := i; j < n+1; j += i {
				table[j] = true
			}
		}

	}
	return primes
}
