package factorial

import (
	"log"
	"math"

	"github.com/erikbryant/util-golang/factorials/bins"
	"github.com/erikbryant/util-golang/factorials/dnc"
	"github.com/erikbryant/util-golang/factorials/moessner"
	"github.com/erikbryant/util-golang/factorials/naive"
	"github.com/erikbryant/util-golang/factorials/swing"
)

// Reduce returns nSmall where nSmall <= n and nSmall! % dnc.Mod == n! % dnc.Mod
func Reduce(n int) int {
	// Idea from: https://euler.stephan-brumme.com/160/
	nDigits := int(math.Log10(float64(n)))
	modDigits := int(math.Log10(float64(dnc.Mod)))

	exp := nDigits - modDigits
	for exp > 0 && n%5 == 0 {
		n /= 5
		exp--
	}

	return n
}

// Factorial runs the given algorithm and displays the results for various values
func Factorial(n int, algorithm string) int {
	var factorial func(int) int

	switch algorithm {
	case "naive":
		factorial = naive.Factorial
	case "dnc":
		factorial = dnc.Factorial
	case "bins":
		factorial = bins.Factorial
	case "swing":
		// Very fast, but memory intensive (requires a list of primes up to n)
		factorial = swing.Factorial
	case "moessner":
		// Nifty algorithm, but not very powerful
		factorial = moessner.Factorial
	default:
		log.Fatal("Not a supported algorithm: ", algorithm)
	}

	return factorial(n)
}
