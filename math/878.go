package math

const MODULO = 1000000007

func nthMagicalNumber(n int, a int, b int) int {
	lcm := FindLCM(a, b)
	repeatedSectionSize := lcm/a + lcm/b - 1
	totalSectionNums := n / repeatedSectionSize
	val := multiply(lcm, totalSectionNums)
	n = n % repeatedSectionSize

	aVal := val
	bVal := val
	for i := 0; i < n; i++ {
		if aVal+a < bVal+b {
			aVal += a
			val = aVal
		} else {
			bVal += b
			val = bVal
		}
	}
	return val % MODULO
}

func multiply(a, b int) int {
	if b == 0 {
		return 0
	}
	return (2*multiply(a, b/2) + a*(b%2)) % MODULO
}
