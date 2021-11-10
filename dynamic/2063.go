package dynamic

import "github.com/gbochora/leetcode/stringsmanipulation"

func countVowels(word string) int64 {
	var answer, sum int64
	for i, c := range word {
		if stringsmanipulation.IsVowel(c) {
			sum += int64(i + 1)
		}
		answer += sum
	}
	return answer
}
