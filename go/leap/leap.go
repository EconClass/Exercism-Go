// Package leap verifies if a given year is a leap year.
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return true
}
