package binarytrees

import (
	"container/list"
	"strconv"
	"strings"
)

//input format: [1,null,2,null,0,3]
func ArrayToBinaryTree(tree string) *TreeNode {
	if len(tree) == 0 {
		return nil
	}
	tree = tree[1 : len(tree)-1]
	parts := strings.Split(tree, ",")
	queue := list.New()
	root := constructTreeNode(parts[0])
	queue.PushBack(root)
	index := 1
	for queue.Len() > 0 && index < len(parts) {
		elem := queue.Front()
		queue.Remove(elem)
		node := (*elem).Value.(*TreeNode)
		if node == nil {
			continue
		}
		node.Left = constructTreeNode(parts[index])
		index++
		if index >= len(parts) {
			break
		}
		node.Right = constructTreeNode(parts[index])
		index++
		queue.PushBack(node.Left)
		queue.PushBack(node.Right)
	}
	return root
}

func constructTreeNode(val string) *TreeNode {
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return nil
	}
	node := new(TreeNode)
	node.Left = nil
	node.Right = nil
	node.Val = valInt
	return node
}
