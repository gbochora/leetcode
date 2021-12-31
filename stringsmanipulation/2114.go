package stringsmanipulation

import "strings"

func mostWordsFound(sentences []string) int {
	max := 0
	for _, s := range sentences {
		words := strings.Split(s, " ")
		if max < len(words) {
			max = len(words)
		}
	}
	return max
}
