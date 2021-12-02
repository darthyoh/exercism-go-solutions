package binarysearch

//SearchInts performs a binary search in a int array
func SearchInts(arr []int, val int, offsets ...int) int {
	if len(offsets) > 1 || len(arr) == 0 {
		return -1
	}

	offset := 0

	if len(offsets) == 1 {
		offset = offsets[0]
	}

	middleValue, index := arr[len(arr)/2], len(arr)/2

	switch {
	case middleValue == val:
		return index + offset
	case middleValue > val:
		return SearchInts(arr[:index], val, offset)
	default:
		return SearchInts(arr[index+1:], val, index+1+offset)
	}
}
