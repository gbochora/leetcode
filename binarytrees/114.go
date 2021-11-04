package binarytrees

func flatten(root *TreeNode) {
	flattenRec(root)
}

func flattenRec(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftTail := flattenRec(root.Left)
	rightTail := flattenRec(root.Right)
	savedRight := root.Right
	root.Right = root.Left
	root.Left = nil
	if leftTail != nil {
		leftTail.Right = savedRight
	} else {
		root.Right = savedRight
	}
	last := rightTail
	if last == nil {
		last = leftTail
	}
	if last == nil {
		last = root
	}
	return last
}
