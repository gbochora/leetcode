package guessnumber

var pick = 10

func guessNumber(n int) int {
	return binarySearch(1, n)
}

func binarySearch(left, right int) int {
	mid := (left + right) / 2
	res := guess(mid)
	if res == 0 {
		return mid
	}
	if res < 0 {
		return binarySearch(left, mid-1)
	} else {
		return binarySearch(mid+1, right)
	}
}

func guess(num int) int {
	return pick - num
}
