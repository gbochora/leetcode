package implementation

func targetIndices(nums []int, target int) []int {
	countLess := 0
	countMore := 0
	for _, v := range nums {
		if v < target {
			countLess++
		}
		if v > target {
			countMore++
		}
	}
	res := make([]int, len(nums)-countLess-countMore)
	for i := range res {
		res[i] = countLess + i
	}
	return res
}
