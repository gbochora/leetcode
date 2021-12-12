package simulation

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	canA := capacityA
	canB := capacityB
	var answ, posA int
	posB := len(plants) - 1

	for posA < posB {
		if plants[posA] > canA {
			answ++
			canA = capacityA
		}

		canA -= plants[posA]
		posA++

		if plants[posB] > canB {
			answ++
			canB = capacityB
		}
		canB -= plants[posB]
		posB--
	}
	if posA == posB && plants[posA] > canA && plants[posB] > canB {
		answ++
	}
	return answ
}
