package util

import (
	"testing"
)

func TestIsPalindromeString(t *testing.T) {
	testCases := []struct {
		c        string
		expected bool
	}{
		{"", true},
		{"w", true},
		{"aba", true},
		{"aab", false},
		{"-22-", true},
	}

	for _, testCase := range testCases {
		answer := IsPalindromeString(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestIsPalindromeInt(t *testing.T) {
	testCases := []struct {
		c        []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 1}, true},
		{[]int{6, 4, 4, 6}, true},
	}

	for _, testCase := range testCases {
		answer := IsPalindromeInt(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		w1       string
		w2       string
		expected bool
	}{
		{"", "", true},
		{"ab", "ba", true},
		{"ab", "ab", true},
		{"ignore", "region", true},
		{"aaaa", "aaa", false},
		{"dog", "gad", false},
	}

	for _, testCase := range testCases {
		answer := IsAnagram(testCase.w1, testCase.w2)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s/%s expected %t, got %t", testCase.w1, testCase.w2, testCase.expected, answer)
		}
	}
}

func TestCryptoquip(t *testing.T) {
	testCases := []struct {
		w1       string
		w2       string
		expected bool
	}{
		{"", "", true},
		{"feel", "felt", false},
		{"keep", "pool", false},
		{"keep", "loot", true},
		{"keep", "kelp", false},
		{"keep", "toot", false},
		{"abcddeeffaa", "12344556611", true},
		{"aaaa", "aaa", false},
		{"dog", "gad", true},
	}

	for _, testCase := range testCases {
		_, answer := Cryptoquip(testCase.w1, testCase.w2)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s/%s expected %t, got %t", testCase.w1, testCase.w2, testCase.expected, answer)
		}
	}
}

func TestIsDigitPermutation(t *testing.T) {
	testCases := []struct {
		c, d     int
		expected bool
	}{
		{-12, -21, true},
		{0, 0, true},
		{1, 1, true},
		{2, 1, false},
		{2, 2, true},
		{9, 99, false},
		{10, 42, false},
		{123, 231, true},
		{212, 122, true},
		{212, 122222, false},
		{1000, 3293, false},
		{2000, 200, false},
		{17526, 23, false},
		{21222, 122, false},
		{87109, 79180, true},
	}

	for _, testCase := range testCases {
		answer := IsDigitPermutation(testCase.c, testCase.d)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d expected %t, got %t", testCase.c, testCase.d, testCase.expected, answer)
		}
	}
}

func TestPartitions(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{-12, 0},
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{9, 30},
	}

	for _, testCase := range testCases {
		answer := Partitions(testCase.c)
		if len(answer) != testCase.expected {
			t.Errorf("ERROR: For %d expected len=%d, got len=%d", testCase.c, testCase.expected, len(answer))
		}
	}
}
