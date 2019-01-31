package isogram

import (
	"regexp"
	"strings"
)

// IsIsogram determines if a word or phrase is an Isogram
// Meaning that every letter is unique
func IsIsogram(text string) bool {
	response := false
	compare := ""

	cleaned := strings.ToLower(text)

	regex := regexp.MustCompile("[^a-zA-Z]+")

	cleaned = regex.ReplaceAllString(cleaned, "")

	for _, letter := range cleaned {
		if !strings.Contains(compare, string(letter)) {
			compare = compare + string(letter)
		}
	}

	if cleaned == compare {
		response = true
	}
	return response
}
