package threesum

import "sort"

func threeSum(nums []int) [][]int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	triplets := make(map[[3]int]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			third := -nums[i] - nums[j]
			if counts[third] == 0 {
				continue
			}
			if nums[i] == 0 && nums[j] == 0 && counts[0] < 3 {
				continue
			}
			if nums[i] == third && counts[third] < 2 {
				continue
			}
			if nums[j] == third && counts[third] < 2 {
				continue
			}
			triple := [3]int{nums[i], nums[j], -nums[i] - nums[j]}
			sort.Ints(triple[:])
			triplets[triple] = true
		}
	}

	res := make([][]int, len(triplets))
	var index int
	for key := range triplets {
		keyCopy := key
		res[index] = keyCopy[:]
		index++
	}
	return res
}
