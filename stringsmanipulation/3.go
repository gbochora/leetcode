package stringsmanipulation

func lengthOfLongestSubstring(s string) int {
	var from, max int
	chars := make(map[byte]bool)
	for to := range s {
		if !chars[s[to]] {
			chars[s[to]] = true
			if max < len(chars) {
				max = len(chars)
			}
			continue
		}

		for s[from] != s[to] {
			delete(chars, s[from])
			from++
		}
		from++
	}
	if max < len(chars) {
		max = len(chars)
	}
	return max
}
