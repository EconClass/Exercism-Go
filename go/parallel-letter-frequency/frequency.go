package letter

import "sync"

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
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
	)
	wg.Add(len(slc))
	for _, in := range slc {
		go func(s string) {
			mutex.Lock()
			for _, r := range s {
				m[r]++
			}
			mutex.Unlock()
			wg.Done()
		}(in)
	}
	wg.Wait()
	return m
}
