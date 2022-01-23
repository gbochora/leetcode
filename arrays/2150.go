package arrays

func findLonely(nums []int) []int {
	counts := make(map[int]int)
	for _, val := range nums {
		counts[val]++
	}
	answ := make([]int, 0)
	for _, val := range nums {
		if counts[val] == 1 && counts[val+1] == 0 && counts[val-1] == 0 {
			answ = append(answ, val)
		}
	}
	return answ
}
