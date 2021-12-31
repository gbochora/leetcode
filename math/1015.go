package math

func smallestRepunitDivByK(k int) int {
	remList := make(map[int]bool)
	length := 1
	rem := 1
	remList[rem] = true
	for rem%k != 0 {
		rem = (rem*10 + 1) % k
		length++
		if remList[rem] {
			return -1
		}
		remList[rem] = true
	}
	return length
}
