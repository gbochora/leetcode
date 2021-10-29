package implementation

func nextBeautifulNumber(n int) int {
	for i := n + 1; ; i++ {
		if isBalanced(i) {
			return i
		}
	}
}

func isBalanced(n int) bool {
	counts := make(map[int]int)
	for n > 0 {
		counts[n%10]++
		n = n / 10
	}
	for key, value := range counts {
		if key != value {
			return false
		}
	}
	return true
}
