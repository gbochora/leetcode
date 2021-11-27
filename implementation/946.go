package implementation

func validateStackSequences(pushed []int, popped []int) bool {
	var pushIndex, popIndex int

	stack := make([]int, 0)
	for popIndex < len(popped) && pushIndex < len(pushed) {
		for pushIndex < len(pushed) && pushed[pushIndex] != popped[popIndex] {
			stack = append(stack, pushed[pushIndex])
			pushIndex++
		}
		if pushIndex < len(pushed) && pushed[pushIndex] == popped[popIndex] {
			pushIndex++
			popIndex++
		}
		for len(stack) > 0 && popped[popIndex] == stack[len(stack)-1] {
			popIndex++
			stack = stack[:len(stack)-1]
		}
	}
	return popIndex == len(popped)
}
