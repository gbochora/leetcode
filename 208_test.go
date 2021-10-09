package main

import "testing"

func TestTrieInsertSearch(t *testing.T) {
	trie := Constructor()
	words := []string{"a", "abc", "aback"}

	for _, w := range words {
		trie.Insert(w)
	}

	for _, w := range words {
		if !trie.Search(w) {
			t.Errorf("can't find word: %v", w)
		}
	}
	notWords := []string{"", "ab", "aba", "abacka"}
	for _, w := range notWords {
		if trie.Search(w) {
			t.Errorf("there should not be  %v, in the trie", w)
		}
	}
}

func TestTrieInsertPrefix(t *testing.T) {
	trie := Constructor()
	words := []string{"dedicate", "dedication", "dead"}

	for _, w := range words {
		trie.Insert(w)
	}

	prefixes := []string{"", "d", "de", "dedi", "dedic", "dedica", "dea", "dead"}
	for _, p := range prefixes {
		if !trie.StartsWith(p) {
			t.Errorf("can't find prefix: %v", p)
		}
	}
	notPrefixes := []string{"a", "ab", "da", "dedicaa", "deade"}
	for _, p := range notPrefixes {
		if trie.Search(p) {
			t.Errorf("there should not be prefix %v, in the trie", p)
		}
	}
}
