package implementation

func isPalindrome(x int) bool {
	rev := reverseNumber(x)
	return rev == x
}

func reverseNumber(x int) int {
	var rev int
	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return rev
}
