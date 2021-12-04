package design

const WordMaxLen = 2000

type StreamChecker struct {
	trie       *Trie
	stream     []byte
	length     int
	pointer    int
	maxWordLen int
}

func ConstructorSC(words []string) StreamChecker {
	var sc StreamChecker
	sc.stream = make([]byte, WordMaxLen)
	sc.pointer = len(sc.stream) - 1
	trie := ConstructorTrie()
	for _, w := range words {
		trie.Insert(reverseString(w))
		if sc.maxWordLen < len(w) {
			sc.maxWordLen = len(w)
		}
	}
	sc.trie = &trie
	return sc
}

func (sc *StreamChecker) Query(letter byte) bool {
	if sc.length < sc.maxWordLen {
		sc.length++
	}
	sc.stream[sc.pointer] = letter
	b := hasWordOnPath(sc.trie, sc.stream, sc.pointer, sc.length)
	sc.pointer--
	if sc.pointer < 0 {
		sc.pointer = len(sc.stream) - 1
	}
	return b
}

func reverseString(s string) string {
	b := make([]byte, len(s))
	for i := range s {
		b[len(s)-i-1] = s[i]
	}
	return string(b)
}

type Trie struct {
	nodes  [26]*Trie
	isWord bool
}

func ConstructorTrie() Trie {
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

func hasWordOnPath(t *Trie, w []byte, pointer, length int) bool {
	if t == nil {
		return false
	}
	if t.isWord || length == 0 {
		return t.isWord
	}
	symbolIndex := w[pointer] - 'a'
	return hasWordOnPath(t.nodes[symbolIndex], w, (pointer+1)%len(w), length-1)
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
