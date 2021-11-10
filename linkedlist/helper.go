package linkedlist

func ArrayToListInt(arr []int) *ListNode {
	var head *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		node := new(ListNode)
		node.Val = arr[i]
		node.Next = head
		head = node
	}
	return head
}
