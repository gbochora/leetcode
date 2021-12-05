package hashtable

import "sort"

func findEvenNumbers(digits []int) []int {
	allDigitsCount := countInts(digits)
	uniqNums := make(map[int]bool)
	for n := 100; n < 1000; n += 2 {
		numDigitsCount := numToDigitsCount(n)
		if hasEnoughDigits(allDigitsCount, numDigitsCount) {
			uniqNums[n] = true
		}
	}

	answ := mapToSlice(uniqNums)
	sort.Ints(answ)
	return answ
}

func countInts(nums []int) map[int]int {
	counts := make(map[int]int)
	for _, n := range nums {
		counts[n]++
	}
	return counts
}

func numToDigitsCount(num int) map[int]int {
	counts := make(map[int]int)
	for num > 0 {
		counts[num%10]++
		num = num / 10
	}
	return counts
}

func hasEnoughDigits(source, num map[int]int) bool {
	for key := range num {
		if num[key] > source[key] {
			return false
		}
	}
	return true
}

func mapToSlice(m map[int]bool) []int {
	slice := make([]int, 0)
	for key := range m {
		slice = append(slice, key)
	}
	return slice
}
