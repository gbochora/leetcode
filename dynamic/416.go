package dynamic

func canPartition(nums []int) bool {
	sums := make(map[int]bool)
	sums[0] = true
	totalSum := getSum(nums)
	if totalSum%2 == 1 {
		return false
	}
	halfSum := totalSum / 2
	for _, v := range nums {
		newSums := make(map[int]bool)
		for key := range sums {
			if key+v == halfSum {
				return true
			}
			if sums[key+v] {
				continue
			}
			newSums[key+v] = true
		}
		for key := range newSums {
			sums[key] = true
		}
	}
	return false
}

func getSum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
