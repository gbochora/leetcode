package implementation

func minStartValue(nums []int) int {
	var min, sum int
	for _, num := range nums {
		sum += num
		if min > sum {
			min = sum
		}
	}
	return 1 - min
}
