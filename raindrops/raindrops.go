package raindrops

import "strconv"

// Convert converts a number to a string
func Convert(num int) string {
	str := ""
	switch {
	case num%3 == 0:
		str = "Pling"
	case num%5 == 0:
		str = "Plang"
	case num%7 == 0:
		str = "Plong"
	default:
		str = strconv.Itoa(num)
	}
	return str
}
