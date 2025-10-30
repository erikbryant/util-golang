package swing

// Prime-Swing factorial
//
// http://www.luschny.de/math/factorial/FastFactorialFunctions.htm
// http://www.luschny.de/math/factorial/SwingIntro.pdf

import (
	"math"

	"github.com/erikbryant/util-golang/factorials/dnc"
	"github.com/erikbryant/util-golang/primey"
)

var (
	// Mod is the global digit mask. Don't change this. Unless you hate yourself.
	Mod = 10000000
)

// multiply returns f, fives where f = p*f*5^fives
func multiply(f, p, fives int) (int, int) {
	for p%5 == 0 {
		fives++
		p /= 5
	}
	f *= p
	f %= Mod
	return f, fives
}

// find returns the index of m in the list of primes or the index of the next higher prime if m is not prime
func find(m int) int {
	i := primey.Index(m)
	if primey.Prime(m) {
		return i
	}
	return i + 1
}

// indices returns the index values for the 4 key variables
func indices(m int) (int, int, int, int) {
	mSqrt := int(math.Sqrt(float64(m)))
	return find(1 + mSqrt), find(1 + m/3), find(1 + m/2), find(1 + m)
}

// swing returns n⎱
func swing(m int) (int, int) {
	if m < 4 {
		return []int{1, 1, 1, 3}[m], 0
	}

	f := 1
	fives := 0

	s, d, e, g := indices(m)

	for _, prime := range primey.Iterr(e, g) {
		f, fives = multiply(f, prime, fives)
	}

	for _, prime := range primey.Iterr(s, d) {
		if (m/prime)&0x01 == 1 {
			f, fives = multiply(f, prime, fives)
		}
	}

	for _, prime := range primey.Iterr(1, s) {
		p, q := 1, m
		for {
			q /= prime
			if q == 0 {
				break
			}
			if q&1 == 1 {
				p *= prime
			}
		}
		if p > 1 {
			f, fives = multiply(f, p, fives)
		}
	}

	return f, fives
}

// factorialOdd returns m and k where 2^?*m*5^k = n!
func factorialOdd(n int) (int, int) {
	if n < 2 {
		return 1, 0
	}

	// f = oddFactorial(n/2, primes)^2 * swing(n, primes)

	f := 1
	fives := 0

	// Highest power of two <= n
	i := int(math.Log2(float64(n)))
	two := int(math.Pow(2, float64(i)))

	for ; two > 0; two /= 2 {
		f *= f
		fives *= 2
		f %= Mod
		fSwing, five := swing(n / two)
		f *= fSwing
		f %= Mod
		fives += five
	}

	return f, fives
}

func Factorial(n int) int {
	twos := dnc.FactorialEven(n)
	f, fives := factorialOdd(n)
	f = dnc.Fix(f, twos, fives)
	return f
}
