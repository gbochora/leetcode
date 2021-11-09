package dynamic

func numTrees(n int) int {
	nums := make([]int, n+1)
	nums[0] = 1
	for i := 1; i <= n; i++ {
		var sum int
		for j := 0; j < i; j++ {
			sum += nums[j] * nums[i-j-1]
		}
		nums[i] = sum
	}
	return nums[n]
}
