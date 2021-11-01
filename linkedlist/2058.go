package linkedlist

func nodesBetweenCriticalPoints(head *ListNode) []int {
	res := []int{-1, -1}
	if head.Next == nil || head.Next.Next == nil {
		return res
	}

	first := head
	mid := head.Next
	last := head.Next.Next
	min := -1
	max := -1
	firstCriticalPoint := -1
	lastCriticalPoint := -1
	index := 1
	for last != nil {
		if first.Val < mid.Val && mid.Val > last.Val ||
			first.Val > mid.Val && mid.Val < last.Val {

			if firstCriticalPoint == -1 {
				firstCriticalPoint = index
				lastCriticalPoint = index
			} else {
				minDistance := index - lastCriticalPoint
				maxDistance := index - firstCriticalPoint
				if max < maxDistance {
					max = maxDistance
				}
				if min == -1 || min > minDistance {
					min = minDistance
				}
				lastCriticalPoint = index
			}
		}
		index++
		first = mid
		mid = last
		last = last.Next
	}
	res[0] = min
	res[1] = max
	return res
}
