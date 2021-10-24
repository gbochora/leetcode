package findwords

import (
	"reflect"
	"testing"
)

func TestFindWords1(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain"}
	got := FindWords(board, words)
	want := []string{"oath", "eat"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindWords2(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'b', 'n'},
		{'o', 't', 'a', 'e'},
		{'a', 'h', 'k', 'r'},
		{'a', 'f', 'l', 'v'}}
	words := []string{"oa", "oaa"}
	got := FindWords(board, words)
	want := []string{"oa", "oaa"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestFindWords3(t *testing.T) {
	board := [][]byte{
		{'a', 'a'}}
	words := []string{"aaa"}
	got := FindWords(board, words)
	want := []string{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
