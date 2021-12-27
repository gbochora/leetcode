package implementation

func getDistances(arr []int) []int64 {
	indexes := make(map[int][]int)
	for i := range arr {
		if indexes[arr[i]] == nil {
			indexes[arr[i]] = make([]int, 0)
		}
		indexes[arr[i]] = append(indexes[arr[i]], i)
	}

	answ := make([]int64, len(arr))
	for i := range arr {
		if indexes[arr[i]] == nil {
			continue
		}
		var sum int64
		for _, v := range indexes[arr[i]] {
			sum += int64(v - i)
		}
		answ[i] = sum
		for j := 1; j < len(indexes[arr[i]]); j++ {
			diff := indexes[arr[i]][j] - indexes[arr[i]][j-1]
			left := j
			right := len(indexes[arr[i]]) - j
			sum += int64(left-right) * int64(diff)
			answ[indexes[arr[i]][j]] = sum
		}
		indexes[arr[i]] = nil
	}
	return answ
}
