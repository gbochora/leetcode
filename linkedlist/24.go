package linkedlist

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := &head
	for *curr != nil && (*curr).Next != nil {
		swapNextTwo(curr)
		curr = &((*curr).Next.Next)
	}

	return head
}

func swapNextTwo(head **ListNode) {
	tmp := *head
	*head = (*head).Next
	tmp.Next = (*head).Next
	(*head).Next = tmp
}
