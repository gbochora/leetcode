package implementation

func getAverages(nums []int, k int) []int {
	avgs := make([]int, len(nums))
	for i := 0; i < k; i++ {
		if i >= len(nums) || len(nums)-i-1 < 0 {
			break
		}
		avgs[i] = -1
		avgs[len(nums)-i-1] = -1
	}
	var sum int64
	for i := 0; i < 2*k+1 && i < len(nums); i++ {
		sum += int64(nums[i])
	}
	for i := k; i < len(nums)-k; i++ {
		avgs[i] = int(sum / int64(2*k+1))
		sum -= int64(nums[i-k])
		if i+k+1 < len(nums) {
			sum += int64(nums[i+k+1])
		}
	}
	return avgs
}
