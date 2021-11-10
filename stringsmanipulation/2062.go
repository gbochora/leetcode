package stringsmanipulation

func countVowelSubstrings(word string) int {
	var counter int
	for i := 0; i < len(word); i++ {
		for j := i + 1; j <= len(word); j++ {
			if isVowelSubstr(word[i:j]) {
				counter++
			}
		}
	}
	return counter
}

func isVowelSubstr(w string) bool {
	vowels := make(map[rune]bool)
	for _, c := range w {
		if !IsVowel(c) {
			return false
		}
		vowels[c] = true
	}
	return len(vowels) == 5
}

func IsVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}
