package raindrops

import "strconv"

// Convert converts a number to a string
func Convert(num int) string {
	str := ""
	if num%3 == 0 {
		str += "Pling"
	}
	if num%5 == 0 {
		str += "Plang"
	}
	if num%7 == 0 {
		str += "Plong"
	}
	if str == "" {
		str = strconv.Itoa(num)
	}
	return str
}
