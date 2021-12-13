package dynamic

func MaxTotalFruits(fruits [][]int, startPos int, k int) int {
	startPos = insertFruitForStart(&fruits, startPos)
	stepsCount := getStepsCount(fruits, startPos)
	leftThenRight := leftThenRight(fruits, stepsCount, startPos, k)
	rightThenLeft := rightThenLeft(fruits, stepsCount, startPos, k)

	if leftThenRight > rightThenLeft {
		return leftThenRight
	}
	return rightThenLeft
}

func leftThenRight(fruits [][]int, stepsCount []int, startPos, k int) int {
	leftPos := startPos
	sum := 0
	for leftPos >= 0 && stepsCount[leftPos] <= k {
		sum += fruits[leftPos][1]
		leftPos--
	}
	leftPos++
	rightPos := startPos
	maxFruits := sum

	for leftPos <= startPos && rightPos < len(fruits)-1 {
		for rightPos < len(fruits)-1 && 2*stepsCount[leftPos]+stepsCount[rightPos+1] <= k {
			rightPos++
			sum += fruits[rightPos][1]
		}
		if sum > maxFruits {
			maxFruits = sum
		}
		sum -= fruits[leftPos][1]
		leftPos++
	}
	return maxFruits
}

func rightThenLeft(fruits [][]int, stepsCount []int, startPos, k int) int {
	sum := 0
	rightPos := startPos
	for rightPos < len(fruits) && stepsCount[rightPos] <= k {
		sum += fruits[rightPos][1]
		rightPos++
	}
	rightPos--
	leftPos := startPos
	maxFruits := sum
	for rightPos >= startPos && leftPos > 0 {
		for leftPos > 0 && 2*stepsCount[rightPos]+stepsCount[leftPos-1] <= k {
			leftPos--
			sum += fruits[leftPos][1]
		}
		if sum > maxFruits {
			maxFruits = sum
		}

		sum -= fruits[rightPos][1]
		rightPos--
	}
	return maxFruits
}

func getStepsCount(fruits [][]int, startPos int) []int {
	steps := make([]int, len(fruits))
	for i := startPos - 1; i >= 0; i-- {
		steps[i] = steps[i+1] - fruits[i][0] + fruits[i+1][0]
	}
	for i := startPos + 1; i < len(fruits); i++ {
		steps[i] = steps[i-1] - fruits[i-1][0] + fruits[i][0]
	}
	return steps
}

func insertFruitForStart(fruits *[][]int, startPos int) int {
	for i := range *fruits {
		if startPos == (*fruits)[i][0] {
			return i
		}
		if startPos < (*fruits)[i][0] {
			*fruits = append(*fruits, []int{0, 0})
			copy((*fruits)[i+1:], (*fruits)[i:])
			(*fruits)[i] = []int{startPos, 0}
			return i
		}
	}
	*fruits = append(*fruits, []int{startPos, 0})
	return len(*fruits) - 1
}
