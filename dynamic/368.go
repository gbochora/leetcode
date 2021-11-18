package dynamic

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dividers := make([][]int, len(nums))
	for i := range nums {
		dividers[i] = make([]int, 0)
		maxDividers := 0
		maxDividersIndex := 0

		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && maxDividers < len(dividers[j]) {
				maxDividers = len(dividers[j])
				maxDividersIndex = j
			}
		}
		if nums[i]%nums[maxDividersIndex] == 0 {
			dividers[i] = append(dividers[i], dividers[maxDividersIndex]...)
		}
		dividers[i] = append(dividers[i], nums[i])
	}

	return getMaxDividersSlice(dividers)
}

func getMaxDividersSlice(dividers [][]int) []int {
	maxIndex := 0
	for i := range dividers {
		if len(dividers[maxIndex]) < len(dividers[i]) {
			maxIndex = i
		}
	}
	return dividers[maxIndex]
}
