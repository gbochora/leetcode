package graphs

import (
	"container/list"
	"math"
)

func maximumDetonation(bombs [][]int) int {
	max := 0
	for i := range bombs {
		val := bfs(bombs, i)
		if val > max {
			max = val
		}
	}
	return max
}

func bfs(bombs [][]int, start int) int {
	queue := list.New()
	visited := make([]bool, len(bombs))
	countBombs := 0
	queue.PushBack(start)
	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		p := elem.Value.(int)
		if visited[p] {
			continue
		}
		visited[p] = true
		countBombs++
		for i := range bombs {
			if !visited[i] && denote(bombs[p], bombs[i]) {
				queue.PushBack(i)
			}
		}
	}
	return countBombs
}

func denote(src, b []int) bool {
	return math.Sqrt(float64(src[0]-b[0])*float64(src[0]-b[0])+float64(src[1]-b[1])*float64(src[1]-b[1])) <= float64(src[2])
}
