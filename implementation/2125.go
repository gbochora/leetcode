package implementation

func numberOfBeams(bank []string) int {
	answ := 0
	i, currRowDevCount := nextRowWithDevices(bank, 0)
	nextNotEmptyRowDevCount := 0
	for ; i < len(bank); i++ {
		i, nextNotEmptyRowDevCount = nextRowWithDevices(bank, i)
		answ += currRowDevCount * nextNotEmptyRowDevCount
		currRowDevCount = nextNotEmptyRowDevCount
		i--
	}
	return answ
}

func nextRowWithDevices(bank []string, from int) (rowIndex, numDevices int) {
	for from < len(bank) && numDevices == 0 {
		numDevices = numDevicesInTheRow(bank[from])
		from++
	}
	return from, numDevices
}

func numDevicesInTheRow(row string) int {
	count := 0
	for i := range row {
		if row[i] == '1' {
			count++
		}
	}
	return count
}
