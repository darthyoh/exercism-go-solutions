package series

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	var found []string

	for i := range make([]int, len(s)+1-n) {
		found = append(found, s[i:i+n])
	}
	return found
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

func First(n int, s string) (first string, ok bool) {
  result := All(n,s)
	if result == nil {
		return "", false
	}

  return result[0], true

}
