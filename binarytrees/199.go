package binarytrees

func rightSideView(root *TreeNode) []int {
	view := make([]int, 0)
	rightSideViewRec(root, 0, &view)
	return view
}

func rightSideViewRec(root *TreeNode, level int, view *[]int) {
	if root == nil {
		return
	}
	if len(*view) <= level {
		*view = append(*view, root.Val)
	}
	rightSideViewRec(root.Right, level+1, view)
	rightSideViewRec(root.Left, level+1, view)
}
