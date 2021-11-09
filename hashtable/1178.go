package hashtable

func findNumOfValidWords(words []string, puzzles []string) []int {
	wordCounts := buildWordCountsMap(words)
	answer := make([]int, len(puzzles))
	for i, p := range puzzles {
		list := make([]int, 0)
		runes := []rune(p)
		generatePuzzleSubsets(runes[1:], 0, getSymbolBit(runes[0]), &list)
		for _, bits := range list {
			answer[i] += wordCounts[bits]
		}
	}
	return answer
}

func buildWordCountsMap(words []string) map[int]int {
	m := make(map[int]int)
	for _, w := range words {
		m[wordToBits(w)]++
	}
	return m
}

func wordToBits(w string) int {
	var bits int
	for _, r := range w {
		bits = bits | getSymbolBit(r)
	}
	return bits
}

func getSymbolBit(r rune) int {
	return 1 << (r - 'a')
}

func generatePuzzleSubsets(p []rune, pos, bits int, list *[]int) {
	if pos >= len(p) {
		*list = append(*list, bits)
		return
	}
	generatePuzzleSubsets(p, pos+1, bits, list)
	generatePuzzleSubsets(p, pos+1, bits|getSymbolBit(p[pos]), list)
}
