package dynamic

func MaximalSquare(matrix [][]byte) int {
	maxHeigh := 0
	for x := 0; x < len(matrix)-maxHeigh; x++ {
		for y := 0; y < len(matrix[x])-maxHeigh; y++ {
			for h := 1; x+h <= len(matrix) && y+h <= len(matrix[x]); h++ {
				if !canBeSquareExpand(matrix, x, y, h) {
					break
				}
				if h > maxHeigh {
					maxHeigh = h
				}
			}
		}
	}
	return maxHeigh * maxHeigh
}

func canBeSquareExpand(matrix [][]byte, x, y, h int) bool {
	for i := 0; i < h; i++ {
		if matrix[x+i][y+h] == '0' {
			return false
		}
		if matrix[x+h][y+i] == '0' {
			return false
		}
	}
	return true
}
