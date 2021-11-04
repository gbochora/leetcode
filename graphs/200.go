package graphs

func numIslands(grid [][]byte) int {
	var num int
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == '1' {
				num++
				dfs(grid, x, y)
			}
		}
	}
	return num
}

func dfs(grid [][]byte, x, y int) {
	grid[x][y] = 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if IsAdjacentCell(dx, dy) && IsInBounds(len(grid), len(grid[0]), x+dx, y+dy) && grid[x+dx][y+dy] == '1' {
				dfs(grid, x+dx, y+dy)
			}
		}
	}
}
