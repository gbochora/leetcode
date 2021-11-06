package binarytrees

func kthSmallest(root *TreeNode, k int) int {
	var counter, answ int
	kthSmallestRec(root, k, &counter, &answ)
	return answ
}

func kthSmallestRec(root *TreeNode, k int, counter, answ *int) {
	if root == nil {
		return
	}
	kthSmallestRec(root.Left, k, counter, answ)
	*counter++
	if *counter == k {
		*answ = root.Val
		return
	}
	kthSmallestRec(root.Right, k, counter, answ)
}
