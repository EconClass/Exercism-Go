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

// ConcurrentFrequency crates a map of unique runes as keys and their respective
// frequencies as values using a concurrent design.
func ConcurrentFrequency(slc []string) FreqMap {
	m := FreqMap{}

	func(s1 string) {
		for _, r := range s1 {
			m[r]++
		}
	}(slc[0])

	func(s2 string) {
		for _, r := range s2 {
			m[r]++
		}
	}(slc[1])

	func(s3 string) {
		for _, r := range s3 {
			m[r]++
		}
	}(slc[2])

	return m
}
