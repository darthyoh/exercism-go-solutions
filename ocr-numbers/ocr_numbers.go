package ocr

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`^ [ _] [ |][ _][ |][ |][ _][ |]`)

var table = map[int]int{
	490: 0,
	288: 1,
	242: 2,
	434: 3,
	312: 4,
	410: 5,
	474: 6,
	290: 7,
	506: 8,
	442: 9,
}

//Recognize ocr digits
func Recognize(str string) []string {
	res := make([]string, 0)
	rows := strings.Split(str, "\n")
	for i := 0; i < len(rows)/4; i++ {
		res = append(res, recognizeRow(rows[i*4+1:i*4+4]))
	}
	return res
}

func recognizeRow(rows []string) string {
	digits := make([]string, len(rows[0])/3)
	for i := 0; i < len(rows[0])/3; i++ {
		for _, row := range rows {
			digits[i] += row[i*3 : i*3+3]
		}
	}
	strs := make([]string, len(digits))
	for i, d := range digits {
		strs[i] = recognizeDigit(d)
	}
	return strings.Join(strs, "")
}

func recognizeDigit(digit string) string {
	if !re.MatchString(digit) {
		return "?"
	}
	sum := 0
	for i, r := range digit {
		if r != 32 {
			sum += int(math.Pow(2, float64(i)))
		}
	}

	if v, ok := table[sum]; ok {
		return fmt.Sprintf("%v", v)
	}

	return "?"
}
