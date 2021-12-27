package sorting

import "sort"

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})
	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := len(merged) - 1
		if merged[last][1] > intervals[i][1] {
			continue
		}
		if merged[last][1] < intervals[i][0] {
			merged = append(merged, intervals[i])
			continue
		}
		merged[last][1] = intervals[i][1]
	}
	return merged
}
