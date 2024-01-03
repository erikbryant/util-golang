package pigLatin

import "strings"

// vowel returns true if the input is a vowel
func vowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

// vowelSound returns true if the string starts with a vowel sound
func vowelSound(s string) bool {
	if len(s) == 0 {
		return false
	}

	if s[0] == 'x' {
		// For example, 'xray'
		return true
	}

	if s[0] == 'y' && len(s) > 1 && !vowel(s[1]) {
		// For example, 'ytterbium'
		return true
	}

	return vowel(s[0])
}

// pigLatin translates a single word to Pig Latin
func pigLatin(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if len(s) == 0 {
		return ""
	}

	// If a word begins with a consonant sound, move it to the end of the word.
	// If a word begins with a vowel sound, do not shift the letters. Note that
	// 'x' and 'y' at the beginning of a word can make vowel sounds.
	// e.g. "xray" -> "xrayway", "yttria" -> "yttriaway"
	if vowelSound(s) {
		return s + "way"
	}

	// We now know the first letter is a consonant sound
	i := 1

	// A 'y' after a consonant cluster makes a vowel sound.
	// e.g. "rhythm" -> "ythmrhay", "my" -> "ymay"
	for i < len(s) && !vowel(s[i]) && s[i] != 'y' {
		i++
	}

	// Rule 3: If a word starts with a consonant sound followed by 'qu', also move 'qu'
	// to the end of the word. e.g. "square" -> "aresquay"
	if i < len(s) && s[i] == 'u' && s[i-1] == 'q' {
		i++
	}

	// Move the consonant cluster to the end
	return s[i:] + s[:i] + "ay"
}

// PigLatin returns the Pig Latin translation of a given word or sentence
func PigLatin(s string) string {
	t := []string{}

	for _, word := range strings.Split(s, " ") {
		t = append(t, pigLatin(word))
	}

	return strings.Join(t, " ")
}
