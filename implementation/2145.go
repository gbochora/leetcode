package implementation

func numberOfArrays(differences []int, lower int, upper int) int {
	val := 0
	min := 0
	max := 0
	for i := range differences {
		val += differences[i]
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	if min < 0 {
		lower -= min
	}
	if max > 0 {
		upper -= max
	}
	if upper < lower {
		return 0
	}
	return upper - lower + 1
}
