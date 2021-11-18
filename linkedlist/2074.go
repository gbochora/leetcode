package linkedlist

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	groupIndex := 2
	node := head
	for node != nil {
		groupLen := getGroupLen(node.Next, groupIndex)
		if groupLen == 0 {
			break
		}
		if groupLen%2 == 0 {
			node = reverseList(node, groupIndex)
		} else {
			for i := 0; i < groupLen; i++ {
				node = node.Next
			}
		}
		groupIndex++
	}

	return head
}

func getGroupLen(head *ListNode, groupIndex int) int {
	groupLen := 0
	for groupLen < groupIndex && head != nil {
		head = head.Next
		groupLen++
	}
	return groupLen
}

func reverseList(head *ListNode, n int) *ListNode {
	node := head.Next
	first := head.Next
	for i := 0; i < n; i++ {
		if node == nil {
			break
		}
		tmp := node.Next
		node.Next = head.Next
		head.Next = node
		node = tmp
	}
	if first != nil {
		first.Next = node
	}
	return first
}
