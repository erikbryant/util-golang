package wordify

import "testing"

func TestPopulateHundreds(t *testing.T) {
	testCases := []struct {
		c        int
		expected string
	}{
		{0, "zero"},
		{1, "one"},
		{20, "twenty"},
		{21, "twenty-one"},
		{90, "ninety"},
		{99, "ninety-nine"},
		{100, "one hundred"},
		{900, "nine hundred"},
		{999, "nine hundred ninety-nine"},
	}

	populateHundreds()

	for _, testCase := range testCases {
		answer := WordHundreds[testCase.c]
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %s, got %s", testCase.c, testCase.expected, answer)
		}
	}
}

func TestStringify(t *testing.T) {
	testCases := []struct {
		c        int
		expected string
	}{
		{-1000, "negative one thousand"},
		{-1, "negative one"},
		{0, "zero"},
		{1, "one"},
		{20, "twenty"},
		{21, "twenty-one"},
		{90, "ninety"},
		{99, "ninety-nine"},
		{100, "one hundred"},
		{900, "nine hundred"},
		{999, "nine hundred ninety-nine"},
		{1000, "one thousand"},
		{1024, "one thousand twenty-four"},
		{1999, "one thousand nine hundred ninety-nine"},
		{10000, "ten thousand"},
		{12024, "twelve thousand twenty-four"},
		{19999, "nineteen thousand nine hundred ninety-nine"},
		{1000000, "one million"},
		{999999999, "nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine"},
		{1000000000, "one billion"},
		{1000000001, "one billion one"},
		{1001000000, "one billion one million"},
		{999999999999, "nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine"},
		{1000000000000, "one trillion"},
		{999999999999999, "nine hundred ninety-nine trillion nine hundred ninety-nine billion nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine"},
		{1000000000000000, "one quadrillion"},
	}

	for _, testCase := range testCases {
		answer := Wordify(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected '%s', got '%s'", testCase.c, testCase.expected, answer)
		}
	}
}
