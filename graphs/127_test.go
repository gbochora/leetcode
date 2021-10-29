package graphs

import "testing"

func TestLadderLength1(t *testing.T) {
	got := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	want := 5
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLadderLength2(t *testing.T) {
	got := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})
	want := 0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
