package math

func minMoves(target int, maxDoubles int) int {
	answ := 0
	for target != 1 && maxDoubles > 0 {
		answ += target % 2
		maxDoubles--
		target /= 2
		answ++
	}

	answ += target - 1
	return answ
}
