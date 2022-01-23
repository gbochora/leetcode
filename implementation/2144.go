package implementation

import "sort"

func minimumCost(cost []int) int {
	answ := 0
	sort.Ints(cost)
	i := len(cost) - 1
	for i > 1 {
		answ += cost[i] + cost[i-1]
		i -= 3
	}
	for i >= 0 {
		answ += cost[i]
		i--
	}
	return answ
}
