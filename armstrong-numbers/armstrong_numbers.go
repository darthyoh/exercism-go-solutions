package armstrong

import (
	"fmt"
	"math"
)

//IsNumber return true if i is an Armstrong Number
func IsNumber(i int) bool {
	sum := 0
	str := fmt.Sprintf("%v", i)

	for _, r := range str {
		sum += int(math.Pow(float64(r-'0'), float64(len(str))))
	}

	return sum == i
}
