package priorityqueue

import (
	"container/heap"

	"github.com/gbochora/leetcode/collections"
)

func minimizedMaximum(n int, quantities []int) int {
	pq := make(collections.PriorityQueue, len(quantities))
	for i, q := range quantities {
		pq[i] = &collections.Item{
			Value:    q,
			Count:    1,
			Priority: q,
			Index:    i,
		}
	}
	heap.Init(&pq)
	n = n - len(quantities)
	for n > 0 {
		item := heap.Pop(&pq).(*collections.Item)
		item.Count++
		item.Priority = item.Value / item.Count
		if item.Value%item.Count > 0 {
			item.Priority++
		}
		heap.Push(&pq, item)
		n--
	}
	maxVal := heap.Pop(&pq).(*collections.Item).Priority

	return maxVal
}
