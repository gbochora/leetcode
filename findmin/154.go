package findmin

func findMin2(nums []int) int {
	min := nums[0]
	findMinBinarySearch2(nums, 0, len(nums)-1, &min)
	return min
}

func findMinBinarySearch2(nums []int, first, last int, min *int) {
	if first > last {
		return
	}
	mid := (first + last) / 2
	if *min > nums[mid] {
		*min = nums[mid]
	}
	if nums[first] == nums[mid] && nums[last] == nums[mid] {
		findMinBinarySearch2(nums, first, mid-1, min)
		findMinBinarySearch2(nums, mid+1, last, min)
	} else if nums[last] == nums[mid] {
		findMinBinarySearch2(nums, first, mid-1, min)
	} else if nums[first] == nums[mid] {
		findMinBinarySearch2(nums, mid+1, last, min)
	} else if nums[first] < nums[mid] && nums[mid] < nums[last] {
		findMinBinarySearch2(nums, first, mid-1, min)
	} else if nums[first] > nums[mid] && nums[mid] < nums[last] {
		findMinBinarySearch2(nums, first, mid-1, min)
	} else if nums[first] == nums[last] {
		findMinBinarySearch2(nums, first, mid-1, min)
		findMinBinarySearch2(nums, mid+1, last, min)
	} else {
		findMinBinarySearch2(nums, mid+1, last, min)
	}
}
