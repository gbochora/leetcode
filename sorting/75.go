package sorting

func sortColors(nums []int) {
	var counts [3]int
	for _, k := range nums {
		counts[k]++
	}
	var pos int
	for i, c := range counts {
		for j := 0; j < c; j++ {
			nums[pos] = i
			pos++
		}
	}
}
