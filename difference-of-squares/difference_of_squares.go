package diffsquares

//SquareOfSum function
func SquareOfSum(n int) int {
	return n * (n + 1) * n * (n + 1) / 4
	// another solution is to use the math.Pow() function
	// but it seems less efficient in benchmark...
}

//SumOfSquares function
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

//Difference function
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
