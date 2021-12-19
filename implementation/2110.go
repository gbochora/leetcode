package implementation

func getDescentPeriods(prices []int) int64 {
	var count int64
	for i := 0; i < len(prices); i++ {
		j := i + 1
		for j < len(prices) && prices[j-1]-prices[j] == 1 {
			j++
		}
		size := int64(j - i)
		count += size * (size + 1) / 2
		i = j - 1
	}
	return count
}
