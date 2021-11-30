package accumulate

//Accumulate simulates a "map" function
func Accumulate(input []string, cb func(string) string) []string {
	ret := make([]string, 0)
	for _, v := range input {
		ret = append(ret, cb(v))
	}
	return ret
}
