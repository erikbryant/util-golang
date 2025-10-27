package moessner

// Moessner addition-only factorial
//
// http://www.luschny.de/math/factorial/FastFactorialFunctions.htm
// http://www.luschny.de/math/factorial/csharp/FactorialAdditiveMoessner.cs.html

var (
	// Mod is the global digit mask. Don't change this. Unless you hate yourself.
	Mod = 10000000
)

// Factorial returns n!
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}

	s := make([]int, n+1)
	s[0] = 1

	for m := 1; m <= n; m++ {
		s[m] = 0
		for k := m; k >= 1; k-- {
			for i := 1; i <= k; i++ {
				s[i] += s[i-1]
			}
		}
	}

	for s[n]%10 == 0 {
		s[n] /= 10
	}

	return s[n] % Mod
}
