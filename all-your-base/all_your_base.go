package allyourbase

import "math"
import "fmt"

//ConvertToBase converts from inputBase to outputBase an inputDigits
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}
	base10input, err := toDecimal(inputBase, inputDigits)
	if err != nil {
		return nil, err
	}

	return reverse(toBase(base10input, outputBase)), nil
}

func toDecimal(inputBase int, inputDigits []int) (int, error) {

	sum := 0
	for i, v := range reverse(inputDigits) {
		if v >= inputBase || v < 0 {
			return 0, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		sum += v * int(math.Pow(float64(inputBase), float64(i)))
	}
	return sum, nil
}

func toBase(input, base int, args ...int) []int {
	if len(args) == 0 {
		args = make([]int, 0)
	}

	args = append(args, input%base)

	input /= base
	if input == 0 {
		return args
	}

	return toBase(input, base, args...)
}

func reverse(inputDigits []int) []int {
	reversed := make([]int, 0)
	for i := len(inputDigits) - 1; i >= 0; i-- {
		reversed = append(reversed, inputDigits[i])
	}
	return reversed
}

/*
41
41/2 => 20
Reste 20
*/
