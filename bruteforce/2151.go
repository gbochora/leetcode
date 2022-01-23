package bruteforce

func maximumGood(statements [][]int) int {
	comb := make([]bool, len(statements))
	var answ int
	maximumGoodRec(statements, comb, 0, &answ)
	return answ
}

func maximumGoodRec(statements [][]int, comb []bool, pos int, max *int) {
	if pos >= len(comb) {
		goodsCount := countGoods(statements, comb)
		if *max < goodsCount {
			*max = goodsCount
		}
		return
	}
	comb[pos] = true
	maximumGoodRec(statements, comb, pos+1, max)
	comb[pos] = false
	maximumGoodRec(statements, comb, pos+1, max)
}

func countGoods(statements [][]int, comb []bool) int {
	count := 0
	for pos, good := range comb {
		if !good {
			continue
		}

		count++
		for i := range statements[pos] {
			if statements[pos][i] == 2 {
				continue
			}
			if (statements[pos][i] == 1 && comb[i]) || (statements[pos][i] == 0 && !comb[i]) {
				continue
			}
			return 0
		}
	}
	return count
}
