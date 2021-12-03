package linkedlist

func oddEvenList(head *ListNode) *ListNode {
	curr := head
	size := listLen(head)
	if size < 3 {
		return head
	}
	tail := getListTail(head)

	for i := 0; i < size/2; i++ {
		tmp := curr.Next
		curr.Next = curr.Next.Next
		tail.Next = tmp
		tail = tail.Next
		tail.Next = nil

		curr = curr.Next
	}
	return head
}

func getListTail(head *ListNode) *ListNode {
	for head != nil && head.Next != nil {
		head = head.Next
	}
	return head
}
