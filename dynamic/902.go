package dynamic

import (
	"math"
	"strconv"
	"strings"
)

func AtMostNGivenDigitSet(digits []string, n int) int {
	max := strings.Split(strconv.Itoa(n), "")
	count := 0
	for i := 1; i < len(max); i++ {
		count += int(math.Pow(float64(len(digits)), float64(i)))
	}
	count += rec(digits, max, 0)
	return count
}

func rec(digits, max []string, index int) int {
	if index >= len(max) {
		return 1
	}
	i := 0
	for i < len(digits) && digits[i] < max[index] {
		i++
	}
	count := i * int(math.Pow(float64(len(digits)), float64((len(max)-index-1))))
	if i < len(digits) && digits[i] == max[index] {
		count += rec(digits, max, index+1)
	}
	return count
}
