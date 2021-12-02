package brackets

var closing = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

//Bracket says if the input string is correctly bracketed !
func Bracket(input string) bool {
	bracketAccumulator := make([]string, 0)

	for _, r := range input {
		switch string(r) {
		case "(", "[", "{":
			bracketAccumulator = append(bracketAccumulator, string(r))
		case ")", "]", "}":
			if len(bracketAccumulator) == 0 || bracketAccumulator[len(bracketAccumulator)-1] != closing[string(r)] {
				return false
			}
			bracketAccumulator = bracketAccumulator[:len(bracketAccumulator)-1]
		}
	}

	return len(bracketAccumulator) == 0
}
