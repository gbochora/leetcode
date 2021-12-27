package linkedlist

func reorderList(head *ListNode) {
	nodes := make([]*ListNode, 0)
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	i := 0
	for ; i < len(nodes)/2; i++ {
		nodes[len(nodes)-i-1].Next = nodes[i].Next
		nodes[i].Next = nodes[len(nodes)-i-1]
	}
	nodes[i].Next = nil
}
