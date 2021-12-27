package sorting

import "sort"

func kClosest(points [][]int, k int) [][]int {
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] < points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	return points[:k]
}
