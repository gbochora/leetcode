package arrays

func rearrangeArray(nums []int) []int {
	answ := make([]int, len(nums))
	var negative, positive int
	i := 0
	for i < len(nums) {
		for nums[positive] < 0 {
			positive++
		}
		for nums[negative] > 0 {
			negative++
		}
		answ[i] = nums[positive]
		i++
		answ[i] = nums[negative]
		i++
		positive++
		negative++
	}
	return answ
}
