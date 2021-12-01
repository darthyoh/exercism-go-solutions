package strain

//Ints type
type Ints []int

//Strings type
type Strings []string

//Lists type
type Lists [][]int

//Keep method for cb predicat
func (ints Ints) Keep(cb func(int) bool) Ints {
	if ints == nil {
		return nil
	}
	newInts := Ints{}

	for _, i := range ints {
		if cb(i) {
			newInts = append(newInts, i)
		}
	}

	return newInts
}

//Discard method for cb predicat
func (ints Ints) Discard(cb func(int) bool) Ints {
	if ints == nil {
		return nil
	}
	newInts := Ints{}

	for _, i := range ints {
		if !cb(i) {
			newInts = append(newInts, i)
		}
	}

	return newInts
}

//Keep method for cb predicat
func (strings Strings) Keep(cb func(string) bool) Strings {
	newStrings := Strings{}

	for _, s := range strings {
		if cb(s) {
			newStrings = append(newStrings, s)
		}
	}

	return newStrings
}

//Keep method for cb predicat
func (lists Lists) Keep(cb func([]int) bool) Lists {
	newLists := Lists{}

	for _, l := range lists {
		if cb(l) {
			newLists = append(newLists, l)
		}
	}

	return newLists
}
