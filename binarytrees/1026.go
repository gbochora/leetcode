package binarytrees

import "math"

func MaxAncestorDiff(root *TreeNode) int {
	max, _, _ := maxAncestorDiffRec(root)
	return max
}

func maxAncestorDiffRec(root *TreeNode) (maxDiff, min, max int) {
	if root.Left == nil || root.Right == nil {
		if root.Left != nil {
			maxDiff, min, max = maxAncestorDiffRec(root.Left)
		} else if root.Right != nil {
			maxDiff, min, max = maxAncestorDiffRec(root.Right)
		} else {
			return 0, root.Val, root.Val
		}
		maxDiff = maxVal(maxDiff, int(math.Abs(float64(root.Val-min))), int(math.Abs(float64(root.Val-max))))
		min = minVal(min, root.Val)
		max = maxVal(max, root.Val)
		return
	}

	leftMaxDiff, leftMin, leftMax := maxAncestorDiffRec(root.Left)
	rightMaxDiff, rightMin, rightMax := maxAncestorDiffRec(root.Right)

	maxDiff = maxVal(leftMaxDiff, rightMaxDiff,
		int(math.Abs(float64(root.Val-leftMax))),
		int(math.Abs(float64(root.Val-leftMin))),
		int(math.Abs(float64(root.Val-rightMax))),
		int(math.Abs(float64(root.Val-rightMin))))

	min = minVal(leftMin, rightMin, root.Val)
	max = maxVal(leftMax, rightMax, root.Val)
	return
}

func minVal(nums ...int) int {
	min := nums[0]
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	return min
}

func maxVal(nums ...int) int {
	max := nums[0]
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	return max
}
