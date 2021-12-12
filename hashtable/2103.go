package hashtable

func countPoints(rings string) int {
	counts := make(map[byte]map[byte]bool)
	for i := range rings {
		if i%2 == 1 {
			continue
		}
		if _, ok := counts[rings[i+1]]; !ok {
			counts[rings[i+1]] = make(map[byte]bool)
		}
		counts[rings[i+1]][rings[i]] = true
	}
	answ := 0
	for key := range counts {
		if len(counts[key]) == 3 {
			answ++
		}
	}
	return answ
}
