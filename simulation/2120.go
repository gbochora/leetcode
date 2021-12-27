package simulation

func executeInstructions(n int, startPos []int, s string) []int {
	answ := make([]int, len(s))
	for i := range s {
		x := startPos[0]
		y := startPos[1]
		j := i
		for j < len(s) {
			switch s[j] {
			case 'L':
				y--
			case 'R':
				y++
			case 'U':
				x--
			case 'D':
				x++
			}
			if !IsInBounds(n, n, x, y) {
				break
			}
			j++
		}
		answ[i] = j - i
	}
	return answ
}

func IsInBounds(w, h, x, y int) bool {
	return x >= 0 && y >= 0 && x < w && y < h
}
