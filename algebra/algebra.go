package algebra

import (
	"math"
)

// NextSquare returns the next highest square and its square root
func NextSquare(n uint) (uint, uint) {
	r := math.Sqrt(float64(n))
	root := uint(r)

	// If n is a square we can save a mul
	if float64(root) == r {
		next := n + root<<1 + 1
		return next, root + 1
	}

	next := root*root + root<<1 + 1
	return next, root + 1
}
