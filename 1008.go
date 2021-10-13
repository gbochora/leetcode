package main

func bstFromPreorder(preorder []int) *TreeNode {
	var root *TreeNode
	for _, val := range preorder {
		insertNode(&root, val)
	}
	return root
}

func insertNode(root **TreeNode, val int) {
	if *root == nil {
		*root = new(TreeNode)
		(*root).Val = val
		return
	}
	if (*root).Val > val {
		insertNode(&(*root).Left, val)
	} else {
		insertNode(&(*root).Right, val)
	}
}
