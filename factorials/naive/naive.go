package naive

// Naive factorial
// Brute force calculates n!

import (
	"log"
)

var (
	// Mod is the global digit mask. Don't change this. Unless you hate yourself.
	Mod = 10000000

	// MaxFives is a value greater than k where k is the largest 5^k factor we expect to encounter
	MaxFives = 16
)

// multiply returns f and k-j where f=[(x * f) % Mod]/(2^k * 5^j)
func multiply(x, f, twos int) (int, int) {
	for twos < MaxFives && x%2 == 0 {
		twos++
		x /= 2
	}

	for x%5 == 0 {
		twos--
		x /= 5
	}

	x %= Mod
	f *= x
	f %= Mod

	return f, twos
}

// fix returns (f*2^twos)%Mod (i.e., it puts back the excess 2's that multiply removed)
func fix(f, twos int) int {
	if twos < 0 || twos > 32 {
		log.Fatal("Twos outside of expected 0-32 range! ", twos)
	}
	return (f << twos) % Mod
}

// Factorial returns [n!/(2^k*5^k)]%Mod
func Factorial(n int) int {
	f := 1
	twos := 0

	for i := 2; i <= n; i++ {
		f, twos = multiply(i, f, twos)
	}
	f = fix(f, twos)

	return f
}
