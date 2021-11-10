package linkedlist

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for i := 1; i < len(lists); i++ {
		lists[0] = mergeTwoLists(lists[0], lists[i])
	}
	return lists[0]
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := l1
	var prev *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev = l1
			l1 = l1.Next
			continue
		}
		if prev != nil {
			prev.Next = l2
		} else {
			head = l2
		}
		tmp := l2.Next
		l2.Next = l1
		prev = l2
		l2 = tmp
	}
	if l2 != nil && prev != nil {
		prev.Next = l2
	}
	return head
}
