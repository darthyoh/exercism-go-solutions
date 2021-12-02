package say

import (
	"fmt"
	"strconv"
	"strings"
)

var table = map[int64]string{
	0:  "zero",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func say99(n int64) (str string) {
	str = ""

	if n >= 100 {
		str = fmt.Sprintf("%v hundred", table[n/100])
		n = n - (100 * (n / 100))
	}
	if n > 0 {
		if str != "" {
			str += " "
		}

		if n < 20 {
			str += table[n]
		} else {
			str += table[10*(n/10)]
			n = n - (10 * (n / 10))
			if n > 0 {
				str += fmt.Sprintf("-%v", table[n])
			}
		}
	}
	return
}

//Say function
func Say(n int64) (string, bool) {

	if n == 0 {
		return "zero", true
	}

	if n < 0 || n > 999999999999 {
		return "", false
	}

	str := fmt.Sprintf("%v", n)
	groups := make([]string, 0)

	temp := str

	for {
		if len(temp) <= 3 {
			groups = append(groups, temp)
			break
		}
		groups = append(groups, temp[len(temp)-3:])
		temp = temp[0 : len(temp)-3]
	}

	says := make([]string, 0)

	for i, v := range groups {
		toAdd := ""
		if j, err := strconv.ParseInt(v, 10, 64); err == nil {
			toAdd = say99(int64(j))

			if toAdd != "" {
				switch i {
				case 1:
					toAdd += " thousand"
				case 2:
					toAdd += " million"
				case 3:
					toAdd += " billion"
				}
				says = append(says, toAdd)
			}

		}
	}

	reversedSays := make([]string, len(says))

	for i, v := range says {
		reversedSays[len(says)-1-i] = v
	}

	return strings.Join(reversedSays, " "), true
}
