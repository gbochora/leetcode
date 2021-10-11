package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	heights := make(map[*TreeNode]int)
	calculateHeightsOfBinaryTree(root, heights)
	return diameterOfBinaryTreeRec(root, heights)
}

func diameterOfBinaryTreeRec(root *TreeNode, heights map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	leftD := diameterOfBinaryTreeRec(root.Left, heights)
	rightD := diameterOfBinaryTreeRec(root.Right, heights)
	maxD := maxInt(leftD, rightD)
	rootD := heights[root.Left] + heights[root.Right]
	return maxInt(maxD, rootD)
}

func calculateHeightsOfBinaryTree(root *TreeNode, heights map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	heightLeft := calculateHeightsOfBinaryTree(root.Left, heights)
	heightRight := calculateHeightsOfBinaryTree(root.Right, heights)
	heights[root] = maxInt(heightLeft, heightRight) + 1
	return heights[root]
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
