package greedy

func minimumBuckets(street string) int {
	symbols := []byte(street)
	res := 0
	for i := range symbols {
		if symbols[i] != 'H' {
			continue
		}

		if i+1 < len(symbols) && symbols[i+1] == '.' {
			res++
			if i+2 < len(symbols) && symbols[i+2] == 'H' {
				symbols[i+2] = 'x'
			}
		} else if i-1 >= 0 && symbols[i-1] == '.' {
			res++
		} else {
			return -1
		}
	}
	return res
}
