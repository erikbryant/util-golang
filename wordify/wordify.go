package wordify

import "strings"

var (
	WordHundreds = []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}
)

// populateHundreds adds the strings from 20-999 to the hundreds slice
func populateHundreds() {
	tens := []string{
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	// 20-99
	for _, ten := range tens {
		WordHundreds = append(WordHundreds, ten)
		for digit := 1; digit <= 9; digit++ {
			WordHundreds = append(WordHundreds, ten+"-"+WordHundreds[digit])
		}
	}

	// 100-999
	for digit := 1; digit <= 9; digit++ {
		h := WordHundreds[digit] + " hundred"
		WordHundreds = append(WordHundreds, h)
		for number := 1; number <= 99; number++ {
			WordHundreds = append(WordHundreds, h+" "+WordHundreds[number])
		}
	}
}

// Wordify returns the words that represent n
func Wordify(n int) string {
	units := []string{
		"",
		"thousand",    // 10^3
		"million",     // 10^6
		"billion",     // 10^9
		"trillion",    // 10^12
		"quadrillion", // 10^15
		"quintillion", // 10^18
		"sextillion",  // 10^21
		"septillion",  // 10^24
		"octillion",   // 10^27
		"nonillion",   // 10^30
		"decillion",   // 10^33
	}

	if len(WordHundreds) < 999 {
		populateHundreds()
	}

	if n == 0 {
		return WordHundreds[0]
	}

	negative := n < 0
	if negative {
		n *= -1
	}

	str := ""
	for _, unit := range units {
		h := n % 1000
		if h != 0 {
			str = WordHundreds[h] + " " + unit + " " + str
		}
		if n == 1 {
			break
		}
		n /= 1000
	}

	if negative {
		str = "negative " + str
	}

	return strings.TrimSpace(str)
}
