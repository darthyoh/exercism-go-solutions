package binarysearchtree

//SearchTreeData is a BST struct
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

//NewBst returns a new BST from an initial Data
func NewBst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

//Insert a data in the tree
func (s *SearchTreeData) Insert(data int) {
	switch {
	case data > s.data:
		if s.right == nil {
			s.right = NewBst(data)
		} else {
			s.right.Insert(data)
		}
	default:
		if s.left == nil {
			s.left = NewBst(data)
		} else {
			s.left.Insert(data)
		}
	}
}

//Map returns ordered values
func (s *SearchTreeData) Map() []int {
	var leftVals, rightVals []int
	if s.left == nil {
		leftVals = []int{}
	} else {
		leftVals = s.left.Map()
	}
	if s.right == nil {
		rightVals = []int{}
	} else {
		rightVals = s.right.Map()
	}
	leftVals = append(leftVals, s.data)
	return append(leftVals, rightVals...)
}

//MapInt maps a function to the tree
func (s *SearchTreeData) MapInt(cb func(int) int) []int {
	ret := make([]int, 0)
	for _, v := range s.Map() {
		ret = append(ret, cb(v))
	}
	return ret
}

//MapString maps a function to the tree
func (s *SearchTreeData) MapString(cb func(int) string) []string {
	ret := make([]string, 0)
	for _, v := range s.Map() {
		ret = append(ret, cb(v))
	}
	return ret
}
