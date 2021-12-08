package binarytrees

import "math"

func findTilt(root *TreeNode) int {
	tilt := 0
	findTiltRec(root, &tilt)
	return tilt
}

func findTiltRec(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}
	leftSum := findTiltRec(root.Left, tilt)
	rightSum := findTiltRec(root.Right, tilt)
	*tilt += int(math.Abs(float64(leftSum - rightSum)))
	return leftSum + rightSum + root.Val
}
