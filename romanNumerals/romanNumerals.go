package romanNumerals

// https://projecteuler.net/about=roman_numerals

// Roman returns the roman numeral expression for n
func Roman(n int) string {
	inflectionPoints := []struct {
		val   int
		roman string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""

	for _, ip := range inflectionPoints {
		for n >= ip.val {
			roman += ip.roman
			n -= ip.val
		}
	}

	return roman
}
