package implementation

func maxSubsequence(nums []int, k int) []int {
	size := len(nums) - k
	for i := 0; i < size; i++ {
		minIndex := 0
		for j := range nums {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums = append(nums[:minIndex], nums[minIndex+1:]...)
	}
	return nums
}
