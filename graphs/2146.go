package graphs

import "sort"

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	savedK := k
	h := len(grid[0])
	w := len(grid)
	visited := make([]bool, w*h)
	distance := make([]int, w*h)
	curr := make([]int, 1)
	curr[0] = pointToNumber(h, start[0], start[1])
	answ := make([][]int, 0)
	lastPos := 0
	currDist := 0
	for len(curr) > 0 && k > 0 {
		numValid := 0
		next := make([]int, 0)
		for _, num := range curr {
			if visited[num] {
				continue
			}
			visited[num] = true
			x, y := numberToPoint(h, num)

			if pricing[0] <= grid[x][y] && pricing[1] >= grid[x][y] {
				numValid++
				answ = append(answ, []int{x, y})
				distance[num] = currDist
			}

			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if IsAdjacentCell(dx, dy) && IsInBounds(w, h, x+dx, y+dy) && grid[x+dx][y+dy] > 0 {
						next = append(next, pointToNumber(h, x+dx, y+dy))
					}
				}
			}
		}

		curr = next
		if numValid <= k {
			k -= numValid
			lastPos += numValid
		} else {
			break
		}
		currDist++
	}
	sort.Slice(answ, func(i int, j int) bool {
		num1 := pointToNumber(h, answ[i][0], answ[i][1])
		num2 := pointToNumber(h, answ[j][0], answ[j][1])
		if distance[num1] != distance[num2] {
			return distance[num1] < distance[num2]
		}

		if grid[answ[i][0]][answ[i][1]] != grid[answ[j][0]][answ[j][1]] {
			return grid[answ[i][0]][answ[i][1]] < grid[answ[j][0]][answ[j][1]]
		}

		if answ[i][0] != answ[j][0] {
			return answ[i][0] < answ[j][0]
		}

		return answ[i][1] < answ[j][1]
	})

	if savedK > len(answ) {
		return answ
	}

	return answ[:savedK]
}
