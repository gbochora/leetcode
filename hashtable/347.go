package hashtable

func topKFrequent(nums []int, k int) []int {
	valueCounts := make(map[int]int)
	for _, val := range nums {
		valueCounts[val]++
	}
	countToVal := make(map[int]map[int]bool)
	for key, val := range valueCounts {
		if _, ok := countToVal[val]; !ok {
			countToVal[val] = make(map[int]bool)
		}
		countToVal[val][key] = true
	}
	kFrequentValues := make([]int, 0)
	for i := len(nums); i > 0; i-- {
		if countToVal[i] == nil {
			continue
		}
		values := countToVal[i]
		for val := range values {
			kFrequentValues = append(kFrequentValues, val)
			if len(kFrequentValues) == k {
				return kFrequentValues
			}
		}
	}
	return kFrequentValues
}
