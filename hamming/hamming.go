package hamming

import "errors"

// Distance counts differences between strings
func Distance(a, b string) (int, error) {
	delta := 0
	if len(a) != len(b) {
		err := errors.New("strings must be of equal length")
		return -1, err
	}
	for i := range a {
		if a[i] != b[i] {
			delta++
		}
	}
	return delta, nil
}
