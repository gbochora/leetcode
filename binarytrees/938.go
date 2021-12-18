package binarytrees

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	return rangeSumBST(root.Right, low, high) + rangeSumBST(root.Left, low, high) + root.Val
}
