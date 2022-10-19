package forth

import (
	"fmt"
	"strconv"
	"strings"
)

func Forth(input []string) ([]int, error) {

	built := make(map[string]string)

	stack := make([]int, 0)

	for _, line := range input {
		line = strings.ToLower(line)

		elems := strings.Fields(line)
		if elems[0] == ":" {
			if _, err := strconv.Atoi(elems[1]); err == nil {
				return nil, fmt.Errorf("illegal operation")
			}
			definition := strings.Join(elems[2:len(elems)-1], " ")
			for k, v := range built {
				definition = strings.ReplaceAll(definition, k, v)
			}
			built[elems[1]] = definition
		} else {
			line = strings.ToLower(line)

			for k, v := range built {
				line = strings.ReplaceAll(line, k, v)
			}

			elems := strings.Fields(line)
			for _, c := range elems {

				if n, err := strconv.Atoi(c); err == nil {
					stack = append(stack, n)
					continue
				} else {
					switch c {
					case "+", "*", "-", "/":
						switch len(stack) {
						case 0:
							return nil, fmt.Errorf("empty stack")
						case 1:
							return nil, fmt.Errorf("only one value on the stack")
						}

						result := 0
						switch c {
						case "+":
							result = stack[len(stack)-2] + stack[len(stack)-1]
						case "-":
							result = stack[len(stack)-2] - stack[len(stack)-1]
						case "*":
							result = stack[len(stack)-2] * stack[len(stack)-1]
						case "/":
							if stack[len(stack)-1] == 0 {
								return nil, fmt.Errorf("divide by zero")
							}
							result = stack[len(stack)-2] / stack[len(stack)-1]
						}
						stack[len(stack)-2] = result
						stack = stack[:len(stack)-1]
						continue
					case "dup":
						if len(stack) == 0 {
							return nil, fmt.Errorf("empty stack")
						}
						stack = append(stack, stack[len(stack)-1])
						continue
					case "drop":
						if len(stack) == 0 {
							return nil, fmt.Errorf("empty stack")
						}
						stack = stack[:len(stack)-1]
						continue
					case "swap":
						switch len(stack) {
						case 0:
							return nil, fmt.Errorf("empty stack")
						case 1:
							return nil, fmt.Errorf("only one value on the stack")
						}
						stack[len(stack)-1], stack[len(stack)-2] = stack[len(stack)-2], stack[len(stack)-1]
						continue
					case "over":
						switch len(stack) {
						case 0:
							return nil, fmt.Errorf("empty stack")
						case 1:
							return nil, fmt.Errorf("only one value on the stack")
						}
						stack = append(stack, stack[len(stack)-2])
						continue
					default:
						return nil, fmt.Errorf("illegal operation")
					}
				}
			}
		}
	}
	return stack, nil
}
