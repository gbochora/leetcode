package tree

func countHighestScoreNodes(parents []int) int {
	tree := buildTreeSlice(parents)
	sizes := make(map[int]int)
	calcTreeSizes(0, tree, sizes)

	scores := make(map[int]int)
	var max, count int

	for i := 0; i < len(parents); i++ {
		scores[i] = calcScore(sizes[tree[i][0]], sizes[tree[i][1]], len(parents)-1-sizes[tree[i][0]]-sizes[tree[i][1]])
		if scores[i] > max {
			max = scores[i]
			count = 1
		} else if scores[i] == max {
			count++
		}
	}

	return count
}

func calcTreeSizes(root int, tree [][]int, sizes map[int]int) int {
	if root == -1 {
		return 0
	}
	sizes[root] = calcTreeSizes(tree[root][0], tree, sizes) + calcTreeSizes(tree[root][1], tree, sizes) + 1
	return sizes[root]
}

func buildTreeSlice(parents []int) [][]int {
	tree := make([][]int, len(parents))
	for i := range tree {
		tree[i] = make([]int, 2)
		tree[i][0] = -1
		tree[i][1] = -1
	}

	for i, p := range parents {
		if p == -1 {
			continue
		}
		if tree[p][0] == -1 {
			tree[p][0] = i
		} else {
			tree[p][1] = i
		}
	}
	return tree
}

func calcScore(sizes ...int) int {
	total := 1
	for _, s := range sizes {
		if s > 0 {
			total *= s
		}
	}
	return total
}
