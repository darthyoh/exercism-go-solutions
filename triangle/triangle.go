// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

// Kind for triangle type
type Kind int

const (
	// NaT not a triangle
	NaT = iota
	// Equ equilateral triangle
	Equ
	// Iso isocele triangle
	Iso
	// Sca scalene triangle
	Sca
)

// KindFromSides function return the kind of triangle
// with a, b and c sides
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || b+c < a || math.IsNaN(a+b+c) || math.IsInf(a+b+c, +1) || math.IsInf(a+b+c, -1) {
		k = NaT
	} else if a == b && a == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
