package wordy

import (
	"strconv"
	"strings"
)

//Answer returns the result of the operation described in question
func Answer(question string) (int, bool) {
	if question[0:8] != "What is " || question[len(question)-1:] != "?" {
		return 0, false
	}

	question = strings.ReplaceAll(question, "multiplied by", "multipliedby")
	question = strings.ReplaceAll(question, "divided by", "dividedby")

	terms := strings.Fields(question[8 : len(question)-1])

	if len(terms)%2 == 0 {
		return 0, false
	}

	for i, term := range terms {
		if i%2 == 0 {
			if _, err := strconv.Atoi(term); err != nil {
				return 0, false
			}
		} else {
			switch term {
			case "plus", "minus", "dividedby", "multipliedby":
				continue
			default:
				return 0, false
			}
		}
	}

	result, _ := strconv.Atoi(terms[0])

	for i := 2; i < len(terms); i += 2 {
		n, _ := strconv.Atoi(terms[i])
		switch terms[i-1] {
		case "plus":
			result += n
		case "minus":
			result -= n
		case "multipliedby":
			result *= n
		case "dividedby":
			result /= n
		}
	}
	return result, true
}
