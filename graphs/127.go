package graphs

import (
	"container/list"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	lexicon := buildLexicon(wordList)
	if !lexicon[endWord] {
		return 0
	}
	distances := make(map[string]int)
	distances[beginWord] = 1
	queue := list.New()
	queue.PushBack(beginWord)

	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		currentWord := (*elem).Value.(string)
		symbols := []byte(currentWord)
		for i := range symbols {
			for c := 'a'; c <= 'z'; c++ {
				symbols[i] = byte(c)
				nextWord := string(symbols)

				if nextWord == endWord {
					return distances[currentWord] + 1
				}
				if lexicon[nextWord] && distances[nextWord] == 0 {
					queue.PushBack(nextWord)
					distances[nextWord] = distances[currentWord] + 1
				}
				symbols[i] = currentWord[i]
			}
		}

	}
	return 0
}

func buildLexicon(wordList []string) map[string]bool {
	lexicon := make(map[string]bool)
	for _, w := range wordList {
		lexicon[w] = true
	}
	return lexicon
}
