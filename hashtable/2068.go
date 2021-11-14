package hashtable

func checkAlmostEquivalent(word1 string, word2 string) bool {
	counts := make(map[rune]int)
	for _, b := range word1 {
		counts[b]++
	}
	for _, b := range word2 {
		counts[b]--
	}

	for _, val := range counts {
		if val < -3 || val > 3 {
			return false
		}
	}
	return true
}
