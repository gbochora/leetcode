package main

func minimumMoves(s string) int {
	var counter int
	for i := 0; i < len(s); i++ {
		if s[i] != 'X' {
			continue
		}
		counter++
		i += 2
	}
	return counter
}
