package stringsmanipulation

func maxPower(s string) int {
	var maxPower, pos int
	for pos < len(s) {
		power := 1
		for pos+1 < len(s) && s[pos] == s[pos+1] {
			pos++
			power++
		}
		if maxPower < power {
			maxPower = power
		}
		pos++
	}
	return maxPower
}
