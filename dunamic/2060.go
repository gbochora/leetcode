package dynamic

import (
	"strconv"
	"unicode"
)

var nums1, nums2 [][]int
var cache map[int64]bool

func possiblyEquals(s1 string, s2 string) bool {
	nums1 = genNumsCombinations([]rune(s1))
	nums2 = genNumsCombinations([]rune(s2))
	cache = make(map[int64]bool)

	return possiblyEqualsRec([]rune(s1), []rune(s2), 0, 0, 0, 0)
}

func possiblyEqualsRec(s1, s2 []rune, i1, i2, preffix1, preffix2 int) bool {
	if b, ok := cache[buildKey(i1, i2, preffix1, preffix2)]; ok {
		return b
	}

	if i1 == len(s1) && i2 == len(s2) && preffix1+preffix2 == 0 {
		return true
	}

	if preffix1 > 0 && preffix2 > 0 {
		return possiblyEqualsRec(s1, s2, i1, i2, preffix1-min(preffix1, preffix2), preffix2-min(preffix1, preffix2))
	}

	if (i1 >= len(s1) && preffix1 == 0) || (i2 >= len(s2) && preffix2 == 0) {
		val := buildKey(i1, i2, preffix1, preffix2)
		cache[val] = false
		return false
	}

	if preffix1 > 0 {
		if unicode.IsDigit(s2[i2]) {
			for i := 0; i < 3; i++ {
				if nums2[i2][i] > 0 {
					if possiblyEqualsRec(s1, s2, i1, i2+i+1, preffix1, nums2[i2][i]) {
						return true
					}
				}
			}
			cache[buildKey(i1, i2, preffix1, preffix2)] = false
			return false
		} else {
			return possiblyEqualsRec(s1, s2, i1, i2+1, preffix1-1, preffix2)
		}
	}

	if preffix2 > 0 {
		if unicode.IsDigit(s1[i1]) {
			for i := 0; i < 3; i++ {
				if nums1[i1][i] > 0 {
					if possiblyEqualsRec(s1, s2, i1+i+1, i2, nums1[i1][i], preffix2) {
						return true
					}
				}
			}
			cache[buildKey(i1, i2, preffix1, preffix2)] = false
			return false
		} else {
			return possiblyEqualsRec(s1, s2, i1+1, i2, preffix1, preffix2-1)
		}
	}

	if unicode.IsLetter(s1[i1]) && unicode.IsLetter(s2[i2]) && s1[i1] != s2[i2] {
		cache[buildKey(i1, i2, preffix1, preffix2)] = false
		return false
	}

	if unicode.IsLetter(s1[i1]) && unicode.IsLetter(s2[i2]) && s1[i1] == s2[i2] {
		return possiblyEqualsRec(s1, s2, i1+1, i2+1, preffix1, preffix1)
	}

	if unicode.IsDigit(s1[i1]) {
		for i := 0; i < 3; i++ {
			if nums1[i1][i] > 0 {
				if possiblyEqualsRec(s1, s2, i1+i+1, i2, nums1[i1][i], preffix2) {
					return true
				}
			}
		}
		cache[buildKey(i1, i2, preffix1, preffix2)] = false
		return false
	}
	if unicode.IsDigit(s2[i2]) {
		for i := 0; i < 3; i++ {
			if nums2[i2][i] > 0 {
				if possiblyEqualsRec(s1, s2, i1, i2+i+1, preffix1, nums2[i2][i]) {
					return true
				}
			}
		}
		cache[buildKey(i1, i2, preffix1, preffix2)] = false
		return false
	}

	cache[buildKey(i1, i2, preffix1, preffix2)] = false
	return false
}

func genNumsCombinations(s []rune) [][]int {
	nums := make([][]int, len(s))
	for i := range s {
		if !unicode.IsDigit(s[i]) {
			continue
		}
		nums[i] = make([]int, 3)
		nums[i][0], _ = strconv.Atoi(string(s[i]))

		if i+1 < len(s) && unicode.IsDigit(s[i+1]) {
			nums[i][1], _ = strconv.Atoi(string(s[i : i+2]))
		}
		if i+2 < len(s) && unicode.IsDigit(s[i+2]) {
			nums[i][2], _ = strconv.Atoi(string(s[i : i+3]))
		}
	}
	return nums
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func buildKey(i1, i2, preffix1, preffix2 int) int64 {
	return int64(preffix1)*10000000 + int64(preffix2)*10000 + int64(i1)*100 + int64(i2)
}
