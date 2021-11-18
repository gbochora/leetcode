package hashtable

func findDisappearedNumbers(nums []int) []int {
	uniqNums := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		curr := nums[i] - 1
		for nums[curr] != 0 {
			tmp := nums[curr] - 1
			nums[curr] = 0
			curr = tmp
			uniqNums++
		}
	}
	rest := make([]int, len(nums)-uniqNums)

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			rest[j] = i + 1
			j++
		}
	}

	return rest
}
