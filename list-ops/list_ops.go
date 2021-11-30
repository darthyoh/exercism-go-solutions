package listops

//import "fmt"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for _, v := range s.Reverse() {
		initial = fn(v, initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	newList := make([]int, 0)
	for _, v := range s {
		if fn(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func (s IntList) Length() int {
	return s.Foldl(func(a, _ int) int {
		return a + 1
	}, 0)
}

func (s IntList) Map(fn func(int) int) IntList {
	newList := make([]int, 0)
	for _, v := range s {
		newList = append(newList, fn(v))
	}
	return newList
}

func (s IntList) Reverse() IntList {
	a := make([]int, len(s))
	copy(a, s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func (s IntList) Append(lst IntList) IntList {
	newS := make([]int, s.Length()+lst.Length())
	copy(newS, s)
	for i, v := range lst {
		newS[s.Length()+i] = v
	}
	return newS
}

func (s IntList) Concat(lists []IntList) IntList {
	if len(lists) == 0 {
		return s
	}

	s = s.Append(lists[0])

	return s.Concat(lists[1:])
}
