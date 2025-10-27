package bins

// Binned Factorial
// We only calculate the last N non-zero digits. Bin the numbers from
// 1..n by their digit length. Calculate the factorial of those
// individual bins.

import (
	"fmt"
	"log"
	"math"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	// Mod is the global digit mask. Don't change this. Unless you hate yourself.
	Mod = 10000000

	// MaxFives is a value greater than k where k is the largest 5^k factor we expect to encounter
	MaxFives = 16
)

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

func fix(f, twos int) int {
	if twos < 0 || twos > 32 {
		log.Fatal("Twos outside of expected 0-32 range! ", twos)
	}
	return (f << twos) % Mod
}

func factorial(n int) int {
	f := 1
	twos := 0

	for i := 2; i <= n; i++ {
		f, twos = multiply(i, f, twos)
	}
	f = fix(f, twos)

	return f
}

var (
	noTensCache = map[string]int{}
)

func factorialNoTens(start, n int) int {
	startEnd := fmt.Sprintf("%d-%d", start, n)
	fCache, ok := noTensCache[startEnd]
	if ok {
		return fCache
	}

	f := 1
	twos := 0

	for i := start; i <= n; i++ {
		if i%10 == 0 {
			continue
		}
		f, twos = multiply(i, f, twos)
	}
	f = fix(f, twos)

	noTensCache[startEnd] = f

	return f
}

func power(base, exp int) int {
	f := 1
	twos := 0

	for i := 1; i <= exp; i++ {
		f, twos = multiply(base, f, twos)
	}
	f = fix(f, twos)

	return f
}

func computeDataset(d Dataset, verbose bool) int {
	pOfP := 1

	if verbose {
		fmt.Printf(" Stage Start     Stage End       Product         Count       Product\n")
	}
	for _, stage := range d.stages {
		rp := factorialNoTens(stage.start, stage.end)
		p := power(rp, stage.count)
		pOfP *= p
		pOfP %= Mod
		if verbose {
			fmt.Printf("%12d  %12d  %12d  %12d  %12d\n", stage.start, stage.end, rp, stage.count, p)
		}
	}

	if d.expected > 0 && pOfP != d.expected {
		fmt.Printf("FAIL!!!!!!! expected: %d  got: %d\n", d.expected, pOfP%100000)
	}
	if verbose {
		fmt.Printf("Product of products: %d  last5: %d\n\n", pOfP, pOfP%100000)
	}

	return pOfP
}

var (
	Elevens = []int{
		1, 1, 11, 101, 1001, 10001, 100001, 1000001, 10000001, 100000001, 1000000001, 10000000001, 100000000001,
	}
	Nines = []int{
		1, 9, 99, 999, 9999, 99999, 999999, 9999999, 99999999, 999999999, 9999999999, 99999999999, 999999999999,
	}
	Offsets = []int{
		0, 10, 110, 1110, 11110, 111110, 1111110, 11111110,
	}
	Expected = []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1946112, 8167808, 3416576,
	}
)

type Stage struct {
	start int
	end   int
	count int
}
type Dataset struct {
	upper    int
	stages   []Stage
	expected int
}

func printDataset(d Dataset) {
	p := message.NewPrinter(language.English)
	fmt.Println(p.Sprintf("%d", d.upper))

	for _, stage := range d.stages {
		fmt.Printf("  %d %d %d\n", stage.start, stage.end, stage.count)
	}

	fmt.Printf("  %d\n\n", d.expected)
}

func makeDataset(upper int) Dataset {
	// The goal is that dataset will look like this:
	// 1,000,000,000,000
	//   1 1 111118
	//   2 9 111117
	//   11 99 111116
	//   101 999 111115
	//   1001 9999 111114
	//   10001 99999 111113
	//   100001 999999 111112
	//   1000001 9999999 111111
	//   -1

	// We count up to upper, but get capped by Mod
	maskDigits := int(math.Log10(float64(Mod))) // + 7 // <===== Set maskDigits to >= 12 and it all works!
	upperDigits := int(math.Log10(float64(upper)))
	limit := min(upperDigits, maskDigits) + 1
	offset := Offsets[max(upperDigits-maskDigits, 0)]

	d := Dataset{upper: upper}
	if upper <= 1000*1000*1000 {
		d.expected = factorial(upper)
	} else {
		d.expected = Expected[upperDigits]
	}

	for r := 0; r < limit; r++ {
		count := (limit - r) + offset
		stage := Stage{start: Elevens[r], end: Nines[r], count: count}
		d.stages = append(d.stages, stage)
	}

	return d
}

func FactorialVerbose(n int) int {
	d := makeDataset(n)
	printDataset(d)
	f := computeDataset(d, true)
	return f
}

func Factorial(n int) int {
	d := makeDataset(n)
	f := computeDataset(d, false)
	return f
}
