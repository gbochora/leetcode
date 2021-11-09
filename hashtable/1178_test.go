package hashtable

import (
	"reflect"
	"testing"
)

func TestFindNumOfValidWords(t *testing.T) {
	words := []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}
	puzzles := []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}

	got := findNumOfValidWords(words, puzzles)
	want := []int{1, 1, 3, 2, 4, 0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindNumOfValidWords2(t *testing.T) {
	words := []string{"apple", "pleas", "please"}
	puzzles := []string{"aelwxyz", "aelpxyz", "aelpsxy", "saelpxy", "xaelpsy"}

	got := findNumOfValidWords(words, puzzles)
	want := []int{0, 1, 3, 2, 0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
