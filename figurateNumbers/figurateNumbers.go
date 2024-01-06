package figuratenumbers

// https://en.wikipedia.org/wiki/Figurate_number
// https://mathworld.wolfram.com/FigurateNumber.html

import (
	"math"

	"github.com/erikbryant/util-golang/util"
)

// IsTriangular returns true if n is a triangular number
func IsTriangular(n int) bool {
	// https://en.wikipedia.org/wiki/Triangular_number
	root := math.Sqrt(float64(8*n + 1))
	return util.IsInt(root)
}

// IsPentagonal returns true if n is a pentagonal number
func IsPentagonal(n int) bool {
	// https://en.wikipedia.org/wiki/Pentagonal_number
	root := math.Sqrt(float64(24*n + 1))
	return util.IsInt((1+root)/6) || util.IsInt((1-root)/6)
}

// IsHexagonal returns true if n is a hexagonal number
func IsHexagonal(n int) bool {
	// https://en.wikipedia.org/wiki/Hexagonal_number
	root := math.Sqrt(float64(8*n + 1))
	return util.IsInt((1+root)/4) || util.IsInt((1-root)/4)
}
