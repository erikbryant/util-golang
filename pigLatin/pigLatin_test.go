package pigLatin

import (
	"testing"
)

func TestVowel(t *testing.T) {
	testCases := []struct {
		c        byte
		expected bool
	}{
		{'a', true},
		{'e', true},
		{'i', true},
		{'o', true},
		{'u', true},
		{'y', false},
		{'z', false},
	}

	for _, testCase := range testCases {
		answer := vowel(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %c expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestVowelSound(t *testing.T) {
	testCases := []struct {
		c        string
		expected bool
	}{
		{"", false},
		{"abc", true},
		{"bcd", false},
		{"y", false},
		{"my", false},
		{"ya", false},
		{"yt", true},
		{"xray", true},
	}

	for _, testCase := range testCases {
		answer := vowelSound(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestPigLatinWord(t *testing.T) {
	testCases := []struct {
		c        string
		expected string
	}{
		// My test cases
		{"", ""},
		{" ", ""},
		{" a", "away"},
		{"a ", "away"},
		{" a ", "away"},
		{"a", "away"},
		{"z", "zay"},
		{"jkl", "jklay"},
		{"ab", "abway"},
		{"ba", "abay"},
		{"hymen", "ymenhay"},
		{"yes", "esyay"},
		{"xray", "xrayway"},
		{"yttria", "yttriaway"},
		{"ytterbium", "ytterbiumway"},
		// Cases from https://exercism.org/tracks/go/exercises/pig-latin
		{"cherry", "errychay"},
		{"pie", "iepay"},
		{"three", "eethray"},
		{"school", "oolschay"},
		{"quiet", "ietquay"},
		{"square", "aresquay"},
		{"the", "ethay"},
		{"quick", "ickquay"},
		{"brown", "ownbray"},
		{"fox", "oxfay"},
		{"rhythm", "ythmrhay"},
		{"my", "ymay"},
		// Cases from https://en.wikipedia.org/wiki/Pig_Latin
		{"pig", "igpay"},
		{"latin", "atinlay"},
		{"banana", "ananabay"},
		{"friends", "iendsfray"},
		{"smile", "ilesmay"},
		{"string", "ingstray"},
		{"eat", "eatway"},
		{"omelet", "omeletway"},
		{"are", "areway"},
	}

	for _, testCase := range testCases {
		answer := pigLatinWord(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s expected %s, got %s", testCase.c, testCase.expected, answer)
		}
	}
}

func TestPigLatin(t *testing.T) {
	testCases := []struct {
		c        string
		expected string
	}{
		{"", ""},
		{"a", "away"},
		{"z", "zay"},
		{"this is a test", "isthay isway away esttay"},
		{"negative one million", "egativenay oneway illionmay"},
	}

	for _, testCase := range testCases {
		answer := PigLatin(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s expected '%s', got '%s'", testCase.c, testCase.expected, answer)
		}
	}
}
