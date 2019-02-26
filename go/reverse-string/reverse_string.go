package reverse

// String takes in a string and returns it in reverse order
func String(input string) string {
	runes := []rune(input)
	for i, val := 0, len(runes)-1; i < val; i, val = i+1, val-1 {
		runes[i], runes[val] = runes[val], runes[i]
	}
	return string(runes)

}
