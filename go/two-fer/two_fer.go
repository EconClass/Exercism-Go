package twofer

import "fmt"

// ShareWith returns the following sentece where X is an input:
// "One for X, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
