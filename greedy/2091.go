package greedy

import "math"

func minimumDeletions(nums []int) int {
	minIndex := 0
	maxIndex := 0
	for i := range nums {
		if nums[minIndex] > nums[i] {
			minIndex = i
		}
		if nums[maxIndex] < nums[i] {
			maxIndex = i
		}
	}

	deletions := int(math.Max(float64(minIndex+1), float64(maxIndex+1)))
	tmp := int(math.Max(float64(len(nums)-minIndex), float64(len(nums)-maxIndex)))
	if deletions > tmp {
		deletions = tmp
	}
	tmp = minIndex + 1 + (len(nums) - maxIndex)
	if deletions > tmp {
		deletions = tmp
	}
	tmp = maxIndex + 1 + (len(nums) - minIndex)
	if deletions > tmp {
		deletions = tmp
	}
	return deletions
}
