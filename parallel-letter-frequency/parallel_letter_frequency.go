package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in given texts and returns this
// data as a FreqMap.
func ConcurrentFrequency(strs []string) FreqMap {
	m := FreqMap{}
	ch := make(chan FreqMap)
	for _, s := range strs {
		go func(s string, ch chan FreqMap) {
			ch <- Frequency(s)
		}(s, ch)
	}
	for i := 0; i < len(strs); i++ {
		ret := <-ch
		for k, v := range ret {
			m[k] += v
		}
	}
	return m
}
