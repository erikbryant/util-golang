package figurate

// https://en.wikipedia.org/wiki/Figurate_number
// https://mathworld.wolfram.com/FigurateNumber.html

import (
	"math"

	"github.com/erikbryant/util-golang/algebra"
)

// Triangular returns the nth triangular number
func Triangular(n int) int {
	return n * (n + 1) / 2
}

// IsTriangular returns true if n is a triangular number
func IsTriangular(n int) bool {
	// https://en.wikipedia.org/wiki/Triangular_number
	root := math.Sqrt(float64(8*n + 1))
	return algebra.IsInt(root)
}

// Pentagonal returns the nth pentagonal number
func Pentagonal(n int) int {
	return n * (3*n - 1) / 2
}

// IsPentagonal returns true if n is a pentagonal number
func IsPentagonal(n int) bool {
	// https://en.wikipedia.org/wiki/Pentagonal_number
	if n == 0 {
		return false
	}
	root := math.Sqrt(float64(24*n + 1))
	return algebra.IsInt((root+1)/6) || algebra.IsInt((-root+1)/6)
}

// Hexagonal returns the nth hexagonal number
func Hexagonal(n int) int {
	return n * (2*n - 1)
}

// IsHexagonal returns true if n is a hexagonal number
func IsHexagonal(n int) bool {
	// https://en.wikipedia.org/wiki/Hexagonal_number
	if n == 0 {
		return false
	}
	root := math.Sqrt(float64(8*n + 1))
	return algebra.IsInt((root+1)/4) || algebra.IsInt((-root+1)/4)
}
