package stack

import "strconv"

func Calculate(s string) int {
	nums := make([]int, 0)
	operators := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if isOper(s[i]) {
			for len(operators) > 0 && priority(s[i]) <= priority(operators[len(operators)-1]) {
				a := nums[len(nums)-2]
				b := nums[len(nums)-1]
				nums[len(nums)-2] = eval(a, b, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
				nums = nums[:len(nums)-1]
			}
			operators = append(operators, s[i])
			continue
		}
		num, end := extractNum(s, i)
		nums = append(nums, num)
		i = end - 1
	}
	for len(operators) > 0 {
		a := nums[len(nums)-2]
		b := nums[len(nums)-1]
		nums[len(nums)-2] = eval(a, b, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
		nums = nums[:len(nums)-1]
	}
	return nums[0]
}

func isOper(b byte) bool {
	return b == '+' || b == '-' || b == '*' || b == '/'
}

func priority(op byte) int {
	if op == '+' || op == '-' {
		return 1
	}
	return 2
}

func eval(a, b int, op byte) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func extractNum(s string, pos int) (num, end int) {
	end = pos + 1
	for end < len(s) && isDigit(s[end]) {
		end++
	}
	num, _ = strconv.Atoi(s[pos:end])
	return
}
