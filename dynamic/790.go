package dynamic

const MODULO = 1000000007

func numTilings(n int) int {
	if n < 2 {
		return 1
	}
	answ := make([]int64, n)
	answ[0] = 1
	answ[1] = 2

	for i := 2; i < n; i++ {
		answ[i] += answ[i-1] + answ[i-2] + 2
		for j := 0; j < i-2; j++ {
			answ[i] = (answ[i] + 2*answ[j]) % MODULO
		}
	}
	return int(answ[n-1])
}
