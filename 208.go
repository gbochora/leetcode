package main

type Trie struct {
	nodes  [26]*Trie
	isWord bool
}

func Constructor() Trie {
	var t Trie
	return t
}

func (trie *Trie) Insert(word string) {
	insertRec(trie, word)
}

func (trie *Trie) Search(word string) bool {
	node := find(trie, word)
	return node != nil && node.isWord
}

func (trie *Trie) StartsWith(prefix string) bool {
	return find(trie, prefix) != nil
}

func find(t *Trie, word string) *Trie {
	if len(word) == 0 || t == nil {
		return t
	}
	symbolIndex := word[0] - 'a'
	return find(t.nodes[symbolIndex], word[1:])
}

func insertRec(t *Trie, word string) {
	if len(word) == 0 {
		t.isWord = true
		return
	}
	symbolIndex := word[0] - 'a'
	if t.nodes[symbolIndex] == nil {
		t.nodes[symbolIndex] = new(Trie)
	}
	insertRec(t.nodes[symbolIndex], word[1:])
}
