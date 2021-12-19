package stringsmanipulation

func firstPalindrome(words []string) string {
	for _, w := range words {
		if isPalindrom(w) {
			return w
		}
	}
	return ""
}

func isPalindrom(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
