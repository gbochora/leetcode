package hashtable

import "strings"

func longestPalindrome(words []string) int {
	counts := make(map[string]int)
	for _, w := range words {
		counts[w]++
	}
	answ := 0
	for w := range counts {
		if w[0] > w[1] {
			continue
		}
		if w[0] == w[1] {
			answ += (counts[w] / 2) * 4
			counts[w] -= (counts[w] / 2) * 2
			continue
		}
		var sb strings.Builder
		sb.WriteByte(w[1])
		sb.WriteByte(w[0])
		reversed := sb.String()
		min := counts[w]
		if min > counts[reversed] {
			min = counts[reversed]
		}

		answ += 4 * min
		counts[w] -= min
		counts[reversed] -= min
	}
	for w := range counts {
		if counts[w] > 0 && w[0] == w[1] {
			answ += 2
			break
		}
	}
	return answ
}
