package sorting

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	pairs := make([][]int, 0)
	minDiff := arr[1] - arr[0] + 1
	for i := 0; i < len(arr)-1; i++ {
		if minDiff < arr[i+1]-arr[i] {
			continue
		}
		if minDiff > arr[i+1]-arr[i] {
			pairs = make([][]int, 0)
			minDiff = arr[i+1] - arr[i]
		}
		pairs = append(pairs, []int{arr[i], arr[i+1]})
	}
	return pairs
}
