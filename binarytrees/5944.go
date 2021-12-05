package binarytrees

func getDirections(root *TreeNode, startValue int, destValue int) string {
	n := treeLength(root)
	startPath := make([]byte, n)
	destPath := make([]byte, n)
	findPathFromRoot(root, startValue, startPath, 0)
	findPathFromRoot(root, destValue, destPath, 0)

	i := 0
	for i = 0; i < len(startPath); i++ {
		if startPath[i] == 0 || destPath[i] == 0 || startPath[i] != destPath[i] {
			break
		}
	}
	answ := make([]byte, 0)
	for j := i; j < len(startPath) && startPath[j] != 0; j++ {
		answ = append(answ, 'U')
	}
	for j := i; j < len(destPath) && destPath[j] != 0; j++ {
		answ = append(answ, destPath[j])
	}
	return string(answ)
}

func treeLength(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return treeLength(root.Left) + treeLength(root.Right) + 1
}

func findPathFromRoot(root *TreeNode, val int, path []byte, steps int) bool {
	if root == nil {
		return false
	}
	if root.Val == val {
		return true
	}
	path[steps] = 'L'
	b := findPathFromRoot(root.Left, val, path, steps+1)
	if b {
		return true
	}
	path[steps] = 'R'
	b = findPathFromRoot(root.Right, val, path, steps+1)
	if b {
		return true
	}
	path[steps] = 0
	return false
}
