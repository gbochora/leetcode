package graphs

import "container/list"

const (
	fresh  int = 1
	rotten int = 2
)

func orangesRotting(grid [][]int) int {
	mins := bfsMinsNum(grid)
	if hasFreshOrange(grid) {
		return -1
	}
	return mins
}

func bfsMinsNum(grid [][]int) int {
	queue := createInitialQueue(grid)
	if queue.Len() == 0 {
		return 0
	}

	visited := make(map[int]bool)
	var mins int
	for queue.Len() > 0 {
		size := queue.Len()
		visitedNewPoints := false
		for i := 0; i < size; i++ {
			point := (*queue.Front()).Value.(int)
			queue.Remove(queue.Front())
			if visited[point] {
				continue
			}
			x, y := numberToPoint(len(grid[0]), point)
			visited[point] = true
			visitedNewPoints = true
			grid[x][y] = rotten
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if IsAdjacentCell(dx, dy) && IsInBounds(len(grid), len(grid[0]), x+dx, y+dy) && grid[x+dx][y+dy] == fresh {
						queue.PushBack(pointToNumber(len(grid[0]), x+dx, y+dy))
					}
				}
			}
		}
		if visitedNewPoints {
			mins++
		}
	}
	return mins - 1
}

func createInitialQueue(grid [][]int) *list.List {
	queue := list.New()

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == rotten {
				queue.PushBack(pointToNumber(len(grid[0]), x, y))
			}
		}
	}
	return queue
}

func hasFreshOrange(grid [][]int) bool {
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == fresh {
				return true
			}
		}
	}
	return false
}
