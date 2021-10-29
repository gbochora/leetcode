package implementation

func missingRolls(rolls []int, mean int, n int) []int {
	sum := mean * (n + len(rolls))
	for _, v := range rolls {
		sum -= v
	}
	if sum < n || sum > n*6 {
		return make([]int, 0)
	}

	missing := make([]int, n)
	for i := range missing {
		missing[i] = 1
	}
	sum -= n
	for i := range missing {
		if sum < 6 {
			missing[i] += sum
			break
		}
		missing[i] += 5
		sum -= 5
	}

	return missing
}
