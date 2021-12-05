package linkedlist

func deleteMiddle(head *ListNode) *ListNode {
	listLen := listLen(head)
	if listLen < 2 {
		return nil
	}

	mid := listLen / 2
	curr := head
	for i := 0; i < mid-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return head
}
