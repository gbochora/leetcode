package dynamic

func uniquePaths(m int, n int) int {
	paths := make([][]int, 2)
	paths[0] = make([]int, n)
	paths[1] = make([]int, n)

	for i := range paths[0] {
		paths[0][i] = 1
	}

	for i := 1; i < m; i++ {
		paths[i%2][0] = paths[(i+1)%2][0]
		for j := 1; j < n; j++ {
			paths[i%2][j] = paths[i%2][j-1] + paths[(i+1)%2][j]
		}
	}
	return paths[(m-1)%2][n-1]
}
