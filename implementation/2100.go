package implementation

func goodDaysToRobBank(security []int, time int) []int {
	answ := make([]int, 0)
	if time == 0 {
		for i := range security {
			answ = append(answ, i)
		}
		return answ
	}

	for start := 0; start < len(security)-2*time; start++ {
		res, pos := isNonIncreasing(security, start, time)
		if !res {
			start = pos
			continue
		}
		res, pos = isNonDecreasing(security, start+time, time)
		if !res {
			continue
		}
		answ = append(answ, start+time)
		for start+time+1+time < len(security) &&
			security[start+time] >= security[start+time+1] &&
			security[start+time+time] <= security[start+time+time+1] {

			answ = append(answ, start+time+1)
			start++
		}
	}
	return answ
}

func isNonIncreasing(nums []int, start, time int) (bool, int) {
	for i := 0; i < time; i++ {
		if nums[start+i] < nums[start+i+1] {
			return false, start + i
		}
	}
	return true, -1
}

func isNonDecreasing(nums []int, start, time int) (bool, int) {
	for i := 0; i < time; i++ {
		if nums[start+i] > nums[start+i+1] {
			return false, start + i
		}
	}
	return true, -1
}
