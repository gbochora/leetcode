package implementation

func maxProduct(nums []int) int {
	answ := maxNum(nums)
	for i := 0; i < len(nums); {
		start, end := nextNonZeroSection(nums, i)
		if start >= len(nums) {
			break
		}
		product := getSectionMaxProduct(nums, start, end)

		if product > answ {
			answ = product
		}
		i = end
	}
	return answ
}

func maxNum(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
func getSectionMaxProduct(nums []int, start, end int) int {
	negativeNumsCount, firstNegativeNumIndex, lastNegativeNumIndex := negativeNumsCount(nums, start, end)
	if negativeNumsCount%2 == 0 {
		return subarrayProduct(nums, start, end)
	}

	max := nums[start]
	beforeFirstNegative := subarrayProduct(nums, start, firstNegativeNumIndex)
	afterFirstNegative := subarrayProduct(nums, firstNegativeNumIndex+1, end)
	beforeLastNegative := subarrayProduct(nums, start, lastNegativeNumIndex)
	afterLastNegative := subarrayProduct(nums, lastNegativeNumIndex+1, end)

	if beforeFirstNegative != 0 && max < beforeFirstNegative {
		max = beforeFirstNegative
	}
	if afterFirstNegative != 0 && max < afterFirstNegative {
		max = afterFirstNegative
	}
	if beforeLastNegative != 0 && max < beforeLastNegative {
		max = beforeLastNegative
	}
	if afterLastNegative != 0 && max < afterLastNegative {
		max = afterLastNegative
	}
	return max
}

func subarrayProduct(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	product := 1
	for j := start; j < end; j++ {
		product *= nums[j]
	}
	return product
}

func negativeNumsCount(nums []int, start, end int) (negativeNumsCount, firstNegativeNumIndex, lastNegativeNumIndex int) {
	negativeNumsCount = 0
	firstNegativeNumIndex = -1
	lastNegativeNumIndex = -1
	for j := start; j < end; j++ {
		if nums[j] > 0 {
			continue
		}
		negativeNumsCount++
		lastNegativeNumIndex = j
		if firstNegativeNumIndex == -1 {
			firstNegativeNumIndex = j
		}
	}
	return
}

func nextNonZeroSection(nums []int, from int) (start, end int) {
	for from < len(nums) && nums[from] == 0 {
		from++
	}

	start = from
	end = from
	for end < len(nums) && nums[end] != 0 {
		end++
	}
	return
}
