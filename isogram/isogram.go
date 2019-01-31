package isogram

import (
	"fmt"
	"regexp"
	"strings"
)

// IsIsogram determines if a word or phrase is an Isogram
// Meaning that every letter is unique
func IsIsogram(text string) bool {
	response := false
	compare := ""

	cleaned := strings.ToLower(text)
	// fmt.Println(cleaned)

	regex := regexp.MustCompile("[^a-zA-Z]+")

	cleaned = regex.ReplaceAllString(cleaned, "")
	// fmt.Println(cleaned)

	// loop through all elements in phrase
	for _, letter := range cleaned {
		// fmt.Println(!strings.Contains(compare, cleaned))
		if strings.Contains(string(letter), cleaned) {
			// fmt.Println(letter)
			compare = compare + string(letter)
			fmt.Println(compare)
		}
		fmt.Println(compare)
	}

	// fmt.Println(cleaned)
	// fmt.Println(compare)

	if cleaned == compare {
		response = true
	}
	return response
}

// func main() {
// 	IsIsogram("isogram")
// }
