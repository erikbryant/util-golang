package romanNumerals

import "strings"

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

// Arabic returns the arabic numeral for the roman numeral
func Arabic(roman string) int {
	arabic := 0

	roman = strings.ToUpper(roman)

	// Add a sentinel so we don't have to keep checking for string overrun
	roman += "_"

	for i := 0; i < len(roman); i++ {
		switch roman[i] {
		case 'M':
			arabic += 1000
		case 'D':
			arabic += 500
		case 'C':
			// Check for the subtractive form CD or CM
			switch roman[i+1] {
			case 'D':
				arabic -= 100
			case 'M':
				arabic -= 100
			default:
				arabic += 100
			}
		case 'L':
			arabic += 50
		case 'X':
			// Check for the subtractive form XL or XC
			switch roman[i+1] {
			case 'L':
				arabic -= 10
			case 'C':
				arabic -= 10
			default:
				arabic += 10
			}
		case 'V':
			arabic += 5
		case 'I':
			// Check for the subtractive form IV or IX
			switch roman[i+1] {
			case 'V':
				arabic -= 1
			case 'X':
				arabic -= 1
			default:
				arabic += 1
			}
		case '_':
			// Do nothing
		}

	}

	return arabic
}
