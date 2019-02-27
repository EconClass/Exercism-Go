package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	trimed := strings.TrimSpace(remark)
	switch {
	case trimed == strings.ToUpper(trimed) && trimed != strings.ToLower(trimed) && trimed[len(trimed)-1:] == "?":
		return "Calm down, I know what I'm doing!"
	case trimed == strings.ToUpper(trimed) && trimed != strings.ToLower(trimed):
		return "Whoa, chill out!"
	case trimed == "":
		return "Fine. Be that way!"
	case trimed[len(trimed)-1:] == "?":
		return "Sure."
	default:
		return "Whatever."
	}
}
