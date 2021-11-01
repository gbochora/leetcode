package hashtable

func minimumOperations(nums []int, start int, goal int) int {
	allNums := make(map[int]bool)
	smallNums := make(map[int]bool)
	smallNums[start] = true
	operations := 0
	for len(smallNums) > 0 {
		operations++
		savedAllNumsLen := len(allNums)
		newSmallNums := generateNewNums(nums, smallNums, allNums)

		if allNums[goal] {
			return operations
		}
		//if there are no new nums
		if len(allNums) == savedAllNumsLen {
			return -1
		}
		smallNums = newSmallNums
	}
	return -1
}

func isSmall(x int) bool {
	return x >= 0 && x <= 1000
}

func generateNewNums(nums []int, smallNums map[int]bool, allNums map[int]bool) map[int]bool {
	newSmallNums := make(map[int]bool)
	for key := range smallNums {
		for _, num := range nums {
			k := key + num
			if isSmall(k) && !allNums[k] {
				newSmallNums[k] = true
			}
			allNums[k] = true
			k = key - num
			if isSmall(k) && !allNums[k] {
				newSmallNums[k] = true
			}
			allNums[k] = true
			k = key ^ num
			if isSmall(k) && !allNums[k] {
				newSmallNums[k] = true
			}
			allNums[k] = true
		}
	}
	return newSmallNums
}
