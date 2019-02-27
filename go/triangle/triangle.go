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

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 || math.IsNaN(a+b+c) || math.IsInf(a+b+c, 0) || a+b < c || a+c < b || b+c < a {
		return NaT
	}

	switch {
	case a == b && b == c && c == a:
		return Equ
	case a == b || b == c || c == a:
		return Iso
	default:
		return Sca
	}
}
