package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if listLen(l1) < listLen(l2) {
		swap(&l1, &l2)
	}

	var saved int
	it1 := l1
	it2 := l2
	for it1 != nil && it2 != nil {
		sum := it1.Val + it2.Val + saved
		saved = sum / 10
		it1.Val = sum % 10
		it1 = it1.Next
		it2 = it2.Next
	}

	for it1 != nil && saved > 0 {
		sum := it1.Val + saved
		saved = sum / 10
		it1.Val = sum % 10
		it1 = it1.Next
	}

	if saved > 0 {
		it1 = l1
		for it1.Next != nil {
			it1 = it1.Next
		}
		it1.Next = new(ListNode)
		it1.Next.Val = saved
	}

	return l1
}

func listLen(l *ListNode) int {
	var size int
	for l != nil {
		size++
		l = l.Next
	}
	return size
}

func swap(l1 **ListNode, l2 **ListNode) {
	tmp := *l1
	*l1 = *l2
	*l2 = tmp
}
