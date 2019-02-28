// Package triangle determines if a shape is a triangle given three sides.
// Then it determines what type of triangle it is.
package triangle

import "math"

// Kind is a string
type Kind string

const (
	// NaT not a triangle
	NaT = "NaT"
	// Equ equilateral
	Equ = "Equ"
	// Iso isosceles
	Iso = "Iso"
	// Sca scalene
	Sca = "Sca"
)

func isNotValid(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return true
	}
	if math.IsNaN(a+b+c) || math.IsInf(a+b+c, 0) {
		return true
	}
	return false
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if isNotValid(a, b, c) {
		return NaT
	}
	switch {
	case a+b < c || a+c < b || b+c < a:
		return NaT
	case a == b && b == c && c == a:
		return Equ
	case a == b || b == c || c == a:
		return Iso
	default:
		return Sca
	}
}
