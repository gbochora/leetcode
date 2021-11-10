package dynamic

import "github.com/gbochora/leetcode/stringsmanipulation"

func countVowels(word string) int64 {
	var answer int64
	sums := make([]int64, len(word))
	for i, c := range word {
		if i > 0 {
			sums[i] = sums[i-1]
		}
		if stringsmanipulation.IsVowel(c) {
			sums[i] += int64(i + 1)
		}
		answer += sums[i]
	}
	return answer
}
