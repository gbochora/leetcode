package math

func FindGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return FindGCD(b, a%b)
}

func FindLCM(a, b int) int {
	return a * b / FindGCD(a, b)
}
