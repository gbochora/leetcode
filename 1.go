package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexes := make([]int, 2)
	counts := make(map[int][]int)
	for i, num := range nums {
		counts[num] = append(counts[num], i)
	}

	for i, num := range nums {
		diff := target - num
		//when x + x = target
		if diff == num && len(counts[diff]) > 1 {
			indexes[0] = counts[diff][0]
			indexes[1] = counts[diff][1]
			break
		}
		//when x + y = target
		if diff != num && len(counts[diff]) > 0 {
			indexes[0] = i
			indexes[1] = counts[diff][0]
			break
		}
	}

	return indexes
}

func main() {
	nums := []int{1, 4, 56, 2, 3, 4}
	fmt.Println(twoSum(nums, 8))
}
