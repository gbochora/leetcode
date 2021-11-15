package design

import "strings"

type CombinationIterator struct {
	characters           string
	combinationLength    int
	combination          []int
	pos                  int
	totalCombinationsNum int
	usedNums             map[int]bool
	combPassed           int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	var it CombinationIterator
	it.characters = characters
	it.combinationLength = combinationLength
	it.combination = make([]int, combinationLength)
	for i := range it.combination {
		it.combination[i] = -1
	}
	it.pos = 0
	it.totalCombinationsNum = fact(len(characters)) / (fact(combinationLength) * fact(len(characters)-combinationLength))
	it.usedNums = make(map[int]bool)
	return it
}

func (it *CombinationIterator) Next() string {
	//undo last number
	if it.pos == it.combinationLength {
		it.pos--
	}
	for it.pos < it.combinationLength {
		//choose next num
		nextNum := it.combination[it.pos] + 1
		if nextNum == 0 && it.pos > 0 {
			nextNum = it.combination[it.pos-1] + 1
		}
		for nextNum < len(it.characters) {
			if !it.usedNums[nextNum] {
				it.usedNums[nextNum] = true
				it.usedNums[it.combination[it.pos]] = false
				it.combination[it.pos] = nextNum
				it.pos++
				break
			}
			nextNum++
		}
		//there is no unused number
		if nextNum == len(it.characters) {
			it.usedNums[it.combination[it.pos]] = false
			it.combination[it.pos] = -1
			it.pos--
		}
	}
	it.combPassed++
	return it.buildCombinationString()
}

func (it *CombinationIterator) HasNext() bool {
	return it.combPassed < it.totalCombinationsNum
}

func (it *CombinationIterator) buildCombinationString() string {
	var sb strings.Builder
	for _, i := range it.combination {
		sb.WriteByte(it.characters[i])
	}
	return sb.String()
}

func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}
