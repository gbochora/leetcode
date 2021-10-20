package main

import (
	"strconv"
	"strings"
)

var nums1MaxValues, nums2MaxValues map[string]int

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	skipCount := len(nums1) + len(nums2) - k
	nums1MaxValues = buildMaxValuesMap(nums1, skipCount)
	nums2MaxValues = buildMaxValuesMap(nums2, skipCount)
	maxNum := make([]int, k)
	cache := make(map[string][]int)

	maxNumberRec(nums1, nums2, 0, 0, skipCount, cache, &maxNum, 0)
	return maxNum
}

func maxNumberRec(nums1, nums2 []int, pos1, pos2, skipCount int, cache map[string][]int, max *[]int, pos int) {
	if pos == len(*max) {
		return
	}
	if cache[buildKey(pos1, pos2, skipCount)] != nil {
		copy((*max)[pos:], cache[buildKey(pos1, pos2, skipCount)])
		return
	}

	key1 := buildKey(pos1, skipCount)
	key2 := buildKey(pos2, skipCount)

	i1, d1 := nextDigit(key1, nums1, nums1MaxValues)
	i2, d2 := nextDigit(key2, nums2, nums2MaxValues)

	if d1 > d2 {
		(*max)[pos] = d1
		maxNumberRec(nums1, nums2, i1+1, pos2, skipCount-i1+pos1, cache, max, pos+1)
	} else if d2 > d1 {
		(*max)[pos] = d2
		maxNumberRec(nums1, nums2, pos1, i2+1, skipCount-i2+pos2, cache, max, pos+1)
	} else {
		(*max)[pos] = d1
		maxNumberRec(nums1, nums2, i1+1, pos2, skipCount-i1+pos1, cache, max, pos+1)
		tmp := make([]int, len(*max)-pos)
		copy(tmp, (*max)[pos:])
		maxNumberRec(nums1, nums2, pos1, i2+1, skipCount-i2+pos2, cache, max, pos+1)

		if compareSlices(tmp, (*max)[pos:]) > 0 {
			copy((*max)[pos:], tmp)
		} else {
			copy(tmp, (*max)[pos:])
		}
		if len(cache) > 100000 {
			return
		}
		cache[buildKey(pos1, pos2, skipCount)] = tmp
	}
}

func compareSlices(s1, s2 []int) int {
	for i := range s1 {
		if s1[i] != s2[i] {
			return s1[i] - s2[i]
		}
	}
	return 0
}

func nextDigit(key string, nums []int, maxValues map[string]int) (index, digit int) {
	if !containsKey(maxValues, key) {
		return -1, -1
	}
	return maxValues[key], nums[maxValues[key]]
}

func containsKey(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}

func buildMaxValuesMap(nums []int, skipCount int) map[string]int {
	maxValues := make(map[string]int)
	for start := 0; start < len(nums); start++ {
		for offset := 0; offset <= skipCount; offset++ {
			maxIndex := start
			for i := start + 1; i <= start+offset && i < len(nums); i++ {
				if nums[maxIndex] < nums[i] {
					maxIndex = i
				}
			}
			maxValues[buildKey(start, offset)] = maxIndex
		}
	}
	return maxValues
}

func buildKey(nums ...int) string {
	var sb strings.Builder
	for i := range nums {
		sb.WriteString(strconv.Itoa(nums[i]) + ":")
	}

	return sb.String()
}
