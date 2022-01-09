package implementation

func maximumInvitations(favorite []int) int {
	guestsNums := make([]int, len(favorite))
	inCycle := make([]bool, len(favorite))
	lastGuest := make([]int, len(favorite))
	visited := make([]bool, len(favorite))
	notValid := make([]bool, len(favorite))
	for i := range favorite {
		if visited[i] {
			continue
		}

		guest := i
		num := 0
		for !visited[guest] {
			visited[guest] = true
			num++
			guest = favorite[guest]
		}
		if inCycle[guest] || notValid[guest] {
			start := i
			for start != guest {
				notValid[start] = true
				start = favorite[start]
			}
			continue
		}

		if guestsNums[guest] > 0 {
			start := i
			for start != guest {
				guestsNums[start] = num + guestsNums[guest]
				lastGuest[start] = lastGuest[guest]
				start = favorite[start]
				num--
			}
		} else if guest == favorite[favorite[guest]] {
			start := i
			for start != guest {
				guestsNums[start] = num
				lastGuest[start] = favorite[guest]
				num--
				start = favorite[start]
			}
			guestsNums[guest] = 2
			guestsNums[favorite[guest]] = 2
			lastGuest[guest] = favorite[guest]
			lastGuest[favorite[guest]] = guest
		} else {
			start := i
			for start != guest {
				notValid[start] = true
				start = favorite[start]
				num--
			}
			inCycle[guest] = true
			guestsNums[guest] = num
			start = favorite[guest]
			for start != guest {
				guestsNums[start] = num
				inCycle[start] = true
				start = favorite[start]
			}
		}
	}
	max := 0
	lastGuestMax := make([]int, len(favorite))
	for i, v := range guestsNums {
		if max < v {
			max = v
		}
		if !inCycle[i] && guestsNums[i] > 0 && v > lastGuestMax[lastGuest[i]] {
			lastGuestMax[lastGuest[i]] = v
		}
	}
	sum := 0
	for i := range favorite {
		if i < favorite[i] && i == favorite[favorite[i]] {
			sum += lastGuestMax[i] + lastGuestMax[favorite[i]] - 2
		}
	}
	if sum > max {
		max = sum
	}
	return max
}
