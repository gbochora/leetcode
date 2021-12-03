package dynamic

func rob(nums []int) int {
	cache := make(map[int]int)
	return robRec(nums, 0, cache)
}

func robRec(nums []int, pos int, cache map[int]int) int {
	if pos >= len(nums) {
		return 0
	}
	if _, ok := cache[pos]; ok {
		return cache[pos]
	}

	firstIncl := robRec(nums, pos+2, cache) + nums[pos]
	firstExcl := robRec(nums, pos+1, cache)
	if firstIncl > firstExcl {
		cache[pos] = firstIncl
	} else {
		cache[pos] = firstExcl
	}
	return cache[pos]
}
