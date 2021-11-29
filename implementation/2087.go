package implementation

import "github.com/gbochora/leetcode/utils"

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	direction := []int{utils.SignInt(homePos[0] - startPos[0]), utils.SignInt(homePos[1] - startPos[1])}

	price := 0
	x := startPos[0]
	for x != homePos[0] {
		x += direction[0]
		price += rowCosts[x]
	}

	y := startPos[1]
	for y != homePos[1] {
		y += direction[1]
		price += colCosts[y]
	}
	return price
}
