package hashtable

import "sort"

func groupAnagrams(strs []string) [][]string {
	anagrams := buildAnagramsClusters(strs)

	groups := make([][]string, len(anagrams))
	var index int
	for _, group := range anagrams {
		groups[index] = group
		index++
	}

	return groups
}

func buildAnagramsClusters(strs []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, word := range strs {
		b := []byte(word)
		sort.Slice(b, func(i int, j int) bool { return b[i] < b[j] })
		key := string(b)
		if _, ok := anagrams[key]; !ok {
			anagrams[key] = make([]string, 0)
		}
		anagrams[key] = append(anagrams[key], word)
	}
	return anagrams
}
