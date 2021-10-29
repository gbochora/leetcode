package binarytrees

import (
	"reflect"
	"testing"
)

func TestHighestScores1(t *testing.T) {
	got := countHighestScoreNodes([]int{-1, 2, 0, 2, 0})
	want := 3
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestHighestScores2(t *testing.T) {
	got := countHighestScoreNodes([]int{-1, 2, 0})
	want := 2
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
