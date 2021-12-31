package stack

func canBeValid(s string, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}
	leftChange := 0
	rightChange := 0
	countLeft := 0
	for i := range s {
		if locked[i] == '0' && s[i] == '(' {
			leftChange++
		}
		if locked[i] == '0' && s[i] == ')' {
			rightChange++
		}
		if s[i] == '(' {
			countLeft++
		} else {
			countLeft--
		}
		if countLeft < 0 {
			rightChange--
			countLeft += 2
		}
		if rightChange < 0 {
			return false
		}
	}
	if countLeft == 0 {
		return true
	}
	leftChange = 0
	rightChange = 0
	countLeft = 0
	for i := range s {
		if locked[len(s)-i-1] == '0' && s[len(s)-i-1] == '(' {
			leftChange++
		}
		if locked[len(s)-i-1] == '0' && s[len(s)-i-1] == ')' {
			rightChange++
		}
		if s[len(s)-i-1] == '(' {
			countLeft++
		} else {
			countLeft--
		}
		if countLeft > 0 {
			leftChange--
			countLeft -= 2
		}
		if leftChange < 0 {
			return false
		}
	}

	return true
}
