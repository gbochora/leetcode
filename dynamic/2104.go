package dynamic

func subArrayRanges(nums []int) int64 {
	var answ int64
	for i := 0; i < len(nums)-1; i++ {
		min := nums[i]
		max := nums[i]
		for j := i; j < len(nums); j++ {
			if min > nums[j] {
				min = nums[j]
			}
			if max < nums[j] {
				max = nums[j]
			}
			answ = answ + int64(max) - int64(min)
		}
	}
	return answ
}
