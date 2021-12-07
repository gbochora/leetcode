package binarytrees

func rob(root *TreeNode) int {
	cache := make(map[*TreeNode]int)

	return robRec(root, cache, false)
}

func robRec(root *TreeNode, cache map[*TreeNode]int, passNode bool) int {
	if root == nil {
		return 0
	}
	if passNode {
		return robRec(root.Left, cache, false) + robRec(root.Right, cache, false)
	}
	if _, ok := cache[root]; ok {
		return cache[root]
	}
	rootInAnsw := robRec(root.Left, cache, true) + robRec(root.Right, cache, true) + root.Val
	rootOutAnsw := robRec(root.Left, cache, false) + robRec(root.Right, cache, false)
	cache[root] = rootInAnsw
	if rootInAnsw < rootOutAnsw {
		cache[root] = rootOutAnsw
	}

	return cache[root]
}
