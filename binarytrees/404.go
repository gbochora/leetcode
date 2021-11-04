package binarytrees

func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeavesRec(root, false)
}

func sumOfLeftLeavesRec(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil && isLeft {
		return root.Val
	}

	return sumOfLeftLeavesRec(root.Left, true) + sumOfLeftLeavesRec(root.Right, false)
}
