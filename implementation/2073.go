package implementation

func timeRequiredToBuy(tickets []int, k int) int {
	first := 0
	time := 0
	for tickets[k] > 0 {
		if tickets[first] > 0 {
			tickets[first]--
			time++
		}
		first = (first + 1) % len(tickets)
	}
	return time
}
