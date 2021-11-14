package binarysearch

import "sort"

func maximumBeauty(items [][]int, queries []int) []int {
	sort.SliceStable(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	maxBeauty := make(map[int]int)
	max := 0
	for i := range items {
		if max < items[i][1] {
			max = items[i][1]
		}
		maxBeauty[items[i][0]] = max
	}

	answer := make([]int, len(queries))
	for i := range queries {
		if _, ok := maxBeauty[queries[i]]; ok {
			answer[i] = maxBeauty[queries[i]]
		} else {
			left := 0
			right := len(items) - 1
			for left < right {
				mid := (left + right) / 2
				if items[mid][0] < queries[i] {
					left = mid + 1
				} else if items[mid][0] > queries[i] {
					right = mid - 1
				}
			}
			if items[left][0] < queries[i] {
				answer[i] = maxBeauty[items[left][0]]
			} else if left-1 >= 0 {
				answer[i] = maxBeauty[items[left-1][0]]
			}
		}
	}

	return answer
}
