package arrays

func countElements(nums []int) int {
	min := nums[0]
	max := nums[0]
	for i := range nums {
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	answ := 0
	for i := range nums {
		if nums[i] > min && nums[i] < max {
			answ++
		}
	}
	return answ
}
