package graphs

import "container/list"

func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))
	queue := list.New()
	queue.PushBack(start)
	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		curr := elem.Value.(int)
		if arr[curr] == 0 {
			return true
		}
		if visited[curr] {
			continue
		}
		visited[curr] = true
		if isValidIndex(curr+arr[curr], len(arr)) {
			queue.PushBack(curr + arr[curr])
		}
		if isValidIndex(curr-arr[curr], len(arr)) {
			queue.PushBack(curr - arr[curr])
		}
	}

	return false
}

func isValidIndex(index, len int) bool {
	return index >= 0 && index < len
}
