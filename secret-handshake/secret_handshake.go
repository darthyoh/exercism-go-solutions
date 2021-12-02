package secret

//Handshake function
func Handshake(n uint) []string {

	table := []string{"wink", "double blink", "close your eyes", "jump"}

	arr := make([]string, 0)

	for k, v := range table {
		if n>>k%2 == 1 {
			arr = append(arr, v)
		}
	}

	if n>>4%2 == 1 {
		tempArr := make([]string, len(arr))

		for i := 0; i < len(arr); i++ {
			tempArr[len(arr)-1-i] = arr[i]
		}
		return tempArr
	}

	return arr
}
