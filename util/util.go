package util

import (
	"sort"
	"strings"
)

// heapPermutation generates a permutation using Heap Algorithm
func heapPermutation(digits []int, size int, c chan []int) {
	// https://www.geeksforgeeks.org/heaps-algorithm-for-generating-permutations/

	if size == 1 {
		var temp []int
		for i := 0; i < len(digits); i++ {
			temp = append(temp, digits[i])
		}
		c <- temp
		return
	}

	for i := 0; i < size; i++ {
		heapPermutation(digits, size-1, c)

		// if size is odd, swap first and last element
		// If size is even, swap ith and last element
		swap := 0
		if size%2 == 0 {
			swap = i
		}
		digits[swap], digits[size-1] = digits[size-1], digits[swap]
	}
}

// MakeDigits generates all permutations of the first n digits
func MakeDigits(n int, c chan []int) {
	// For example:
	//
	//	n=2 [1 2] [2 1]
	//	n=3 [1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]

	defer close(c)

	var digits []int
	for i := 1; i <= n; i++ {
		digits = append(digits, i)
	}

	heapPermutation(digits, len(digits), c)
}

// IsPalindromeString returns true if the string is a palindrome
func IsPalindromeString(p string) bool {
	head := 0
	tail := len(p) - 1

	for head < tail {
		if p[head] != p[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}

// IsPalindromeInt returns true if the digits of p are a palindrome
func IsPalindromeInt(p []int) bool {
	head := 0
	tail := len(p) - 1

	for head < tail {
		if p[head] != p[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}

// IsAnagram returns true if w1 and w2 are anagrams of each other
func IsAnagram(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	for _, c := range w1 {
		w2 = strings.Replace(w2, string(c), "", 1)
	}

	return w2 == ""
}

// Cryptoquip returns true if the letter arrangements are similar; e.g., KEEP and LOOT
func Cryptoquip(w1, w2 string) (map[byte]byte, bool) {
	if len(w1) != len(w2) {
		return nil, false
	}

	substitutions := make(map[byte]byte)

	for i := 0; i < len(w1); i++ {
		if val, ok := substitutions[w1[i]]; ok {
			if val != w2[i] {
				return nil, false
			}
			continue
		}
		if val, ok := substitutions[w2[i]]; ok {
			if val != w1[i] {
				return nil, false
			}
			continue
		}
		substitutions[w1[i]] = w2[i]
		substitutions[w2[i]] = w1[i]
	}

	return substitutions, true
}

// IsDigitPermutation returns whether the two numbers are digit permutations of each other
func IsDigitPermutation(a, b int) bool {
	// Take the absolute value
	a = max(a, -1*a)
	b = max(b, -1*b)

	digits := map[int]int{}

	for a > 0 && b > 0 {
		r := a % 10
		digits[r]++
		a /= 10

		r = b % 10
		digits[r]--
		b /= 10
	}

	if a != b {
		// a and b were not the same length
		return false
	}

	for _, val := range digits {
		if val != 0 {
			return false
		}
	}

	return true
}

// Partitions returns all integer partitions of n
func Partitions(n int) [][]int {
	partitions := [][]int{}

	a := make([]int, n)
	k := 1
	a[1] = n
	for k != 0 {
		x := a[k-1] + 1
		y := a[k] - 1
		k -= 1
		for x <= y {
			a[k] = x
			y -= x
			k += 1
		}
		a[k] = x + y
		c := append([]int{}, a[0:k+1]...)
		sort.Sort(sort.Reverse(sort.IntSlice(c)))
		partitions = append(partitions, c)
	}

	return partitions
}
