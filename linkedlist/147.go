package linkedlist

func InsertionSortList(head *ListNode) *ListNode {
	curr := head

	for curr != nil {
		for curr.Next != nil && curr.Val <= curr.Next.Val {
			curr = curr.Next
		}
		if curr.Next == nil {
			break
		}
		node := curr.Next
		curr.Next = curr.Next.Next
		insertInSortedList(&head, node)
	}

	return head
}

func insertInSortedList(head **ListNode, node *ListNode) {
	if head == nil {
		*head = node
		node.Next = nil
		return
	}
	if (*head).Val > node.Val {
		node.Next = (*head)
		*head = node
		return
	}
	curr := *head
	for curr.Next != nil && curr.Next.Val < node.Val {
		curr = curr.Next
	}
	node.Next = curr.Next
	curr.Next = node
}
