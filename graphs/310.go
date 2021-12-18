package graphs

func FindMinHeightTrees(n int, edges [][]int) []int {
	graph := buildGraph(n, edges)
	excluded := make([]bool, n)
	minHeightTrees := make([]int, 0)
	minHeight := n
	for node := range graph {
		if excluded[node] {
			continue
		}
		height, _ := getTreeHeightAndLeaves(graph, node, minHeight)
		if minHeight == height {
			minHeightTrees = append(minHeightTrees, node)
		}
		if minHeight > height {
			minHeight = height
			minHeightTrees = make([]int, 0)
			minHeightTrees = append(minHeightTrees, node)
		}
	}
	return minHeightTrees
}

func getTreeHeightAndLeaves(graph [][]int, root, minHeight int) (int, []int) {
	visited := make([]bool, len(graph))
	level := make([]int, 0)
	level = append(level, root)
	heigh := 0
	for {
		heigh++
		if heigh > minHeight {
			break
		}
		nextLevel := make([]int, 0)
		for _, node := range level {
			visited[node] = true
			for _, child := range graph[node] {
				if visited[child] {
					continue
				}
				nextLevel = append(nextLevel, child)
			}
		}
		if len(nextLevel) == 0 {
			break
		}
		level = nextLevel
	}
	return heigh, level
}

func buildGraph(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for i := range edges {
		if graph[edges[i][0]] == nil {
			graph[edges[i][0]] = make([]int, 0)
		}
		if graph[edges[i][1]] == nil {
			graph[edges[i][1]] = make([]int, 0)
		}
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return graph
}
