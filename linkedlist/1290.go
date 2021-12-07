package linkedlist

import "math"

func getDecimalValue(head *ListNode) int {
	size := listLen(head)
	decimalVal := 0
	for head != nil {
		size--
		decimalVal += head.Val * int(math.Pow(float64(2), float64(size)))
		head = head.Next
	}
	return decimalVal
}
