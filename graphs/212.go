package graphs

import "github.com/gbochora/leetcode/trie"

func FindWords(board [][]byte, words []string) []string {

	trie := buildTrie(words)
	list := dfsFindWords(board, trie)

	return list
}

func dfsFindWords(board [][]byte, trie *trie.Trie) []string {
	generatedWords := make(map[string]bool)
	for x := range board {
		for y := range board[0] {
			visited := createVisitedGrid(len(board), len(board[0]))
			visited[x][y] = true
			dfsFromSpecificCell(board, x, y, string(board[x][y]), visited, trie, generatedWords)
		}
	}
	list := make([]string, 0, len(generatedWords))
	for w := range generatedWords {
		list = append(list, w)
	}
	return list
}

func dfsFromSpecificCell(board [][]byte, x, y int, prefix string, visited [][]bool, trie *trie.Trie, generatedWords map[string]bool) {
	if !trie.StartsWith(prefix) {
		return
	}
	if trie.Search(prefix) {
		generatedWords[prefix] = true
	}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if IsAdjacentCell(dx, dy) && IsInBounds(board, x+dx, y+dy) && !visited[x+dx][y+dy] {
				visited[x+dx][y+dy] = true
				dfsFromSpecificCell(board, x+dx, y+dy, prefix+string(board[x+dx][y+dy]), visited, trie, generatedWords)
				visited[x+dx][y+dy] = false
			}
		}
	}
}

func createVisitedGrid(w, h int) [][]bool {
	visited := make([][]bool, w)
	for i := range visited {
		visited[i] = make([]bool, h)
	}
	return visited
}

func buildTrie(words []string) *trie.Trie {
	trie := trie.Constructor()
	for _, w := range words {
		trie.Insert(w)
	}
	return &trie
}

func IsAdjacentCell(dx, dy int) bool {
	return dx != dy && dx*dy == 0
}

func IsInBounds(board [][]byte, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(board) && y < len(board[0])
}
