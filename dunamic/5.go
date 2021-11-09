package dynamic

func longestPalindrome(s string) string {
	evenFrom, evenMaxSize := findLongestPalindrome(s, 2, 2)
	oddFrom, oddMaxSize := findLongestPalindrome(s, 3, 2)
	if evenMaxSize > oddMaxSize {
		return s[evenFrom : evenFrom+evenMaxSize]
	}
	return s[oddFrom : oddFrom+oddMaxSize]
}

func findLongestPalindrome(s string, initSize, iter int) (from, maxSize int) {
	maxSize = 1
	last := make([]bool, len(s))
	for i := range last {
		last[i] = true
	}
	next := make([]bool, len(s))

	for i := initSize; i <= len(s); i += iter {
		var palindromeFound bool
		for j := 0; j <= len(s)-i; j++ {
			next[j] = last[j+1] && s[j] == s[j+i-1]
			if next[j] {
				palindromeFound = true
				from = j
			}
		}
		if !palindromeFound {
			break
		}
		maxSize = i
		copy(last, next)
	}
	return from, maxSize
}
