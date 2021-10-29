package implementation

import (
	"strings"
	"unicode"
)

func countValidWords(sentence string) int {
	words := strings.Fields(sentence)
	var count int
	for _, w := range words {
		if isValidWord(w) {
			count++
		}
	}
	return count
}

func isValidWord(w string) bool {
	var numHypens int
	runes := []rune(w)
	for i, r := range runes {
		if unicode.IsDigit(r) {
			return false
		}
		if r == '-' {
			numHypens++
			if i == 0 || i == len(w)-1 || numHypens > 1 {
				return false
			}
			if !unicode.IsLetter(runes[i-1]) || !unicode.IsLetter(runes[i+1]) {
				return false
			}
		}

		if (r == '!' || r == '.' || r == ',') && i != len(w)-1 {
			return false
		}

	}
	return true
}
