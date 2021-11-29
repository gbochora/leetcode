package hashtable

func countWords(words1 []string, words2 []string) int {
	counts1 := getCountMap(words1)
	counts2 := getCountMap(words2)

	res := 0
	for key := range counts1 {
		if counts1[key] == 1 && counts2[key] == 1 {
			res++
		}
	}
	return res
}

func getCountMap(words []string) map[string]int {
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}
	return counts
}
