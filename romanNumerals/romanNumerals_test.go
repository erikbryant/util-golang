package romanNumerals

import "testing"

func TestRomanArabic(t *testing.T) {
	testCases := []struct {
		n        int
		expected string
	}{
		{0, ""},
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{12, "XII"},
		{19, "XIX"},
		{29, "XXIX"},
		{39, "XXXIX"},
		{43, "XLIII"},
		{50, "L"},
		{77, "LXXVII"},
		{87, "LXXXVII"},
		{100, "C"},
		{500, "D"},
		{999, "CMXCIX"},
		{1000, "M"},
		{2000, "MM"},
		{2001, "MMI"},
		{9000, "MMMMMMMMM"},
		{9999, "MMMMMMMMMCMXCIX"},
	}

	for _, testCase := range testCases {
		answer := Roman(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %s, got %s", testCase.n, testCase.expected, answer)
		}
		answer2 := Arabic(answer)
		if answer2 != testCase.n {
			t.Errorf("ERROR: For %s expected %d, got %d", answer, testCase.n, answer2)
		}
	}
}
