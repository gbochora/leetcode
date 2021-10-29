package backtracking

const MARKED = 'x'

func permute(nums []int) [][]int {
	permList := make([][]int, 0)
	perm := make([]int, 0)
	permuteRec(nums, perm, &permList)
	return permList
}

func permuteRec(nums []int, perm []int, permList *[][]int) {
	if len(perm) == len(nums) {
		final := make([]int, len(perm))
		copy(final, perm)
		*permList = append(*permList, final)
		return
	}

	for i := range nums {
		num := nums[i]
		if num == MARKED {
			continue
		}
		perm = append(perm, num)
		nums[i] = MARKED //mark as visited
		permuteRec(nums, perm, permList)
		perm = perm[:len(perm)-1]
		nums[i] = num //unmake previous decision
	}
}
