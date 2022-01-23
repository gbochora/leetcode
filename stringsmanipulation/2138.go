package stringsmanipulation

import "strings"

func divideString(s string, k int, fill byte) []string {
	answ := make([]string, 0)
	i := 0
	for ; i+k < len(s); i += k {
		answ = append(answ, s[i:i+k])
	}
	answ = append(answ, s[i:])

	rem := len(s) % k
	var sb strings.Builder
	sb.WriteString(answ[len(answ)-1])
	if rem != 0 {
		for i := 0; i < k-rem; i++ {
			sb.WriteByte(fill)
		}
		answ[len(answ)-1] = sb.String()
	}

	return answ
}
