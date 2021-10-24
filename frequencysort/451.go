package frequencysort

func frequencySort(s string) string {
	runeToFreqMap := buildRuneToFreqMap(s)
	freqToRunesMap := buildFreqTuRunesMap(runeToFreqMap)

	res := make([]rune, len(s))
	var pos int
	for i := len(s); i > 0; i-- {
		if slice, ok := freqToRunesMap[i]; ok {
			for _, r := range slice {
				addRunes(res, pos, r, i)
				pos += i
			}
		}
	}

	return string(res)
}

func buildFreqTuRunesMap(runeToFreqMap map[rune]int) map[int][]rune {
	m := make(map[int][]rune)
	for key, val := range runeToFreqMap {
		if _, ok := m[val]; ok {
			m[val] = append(m[val], key)
		} else {
			m[val] = make([]rune, 0)
			m[val] = append(m[val], key)
		}
	}
	return m
}

func buildRuneToFreqMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	return m
}

func addRunes(res []rune, pos int, r rune, n int) {
	for i := 0; i < n; i++ {
		res[pos+i] = r
	}
}
