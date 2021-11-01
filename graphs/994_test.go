package graphs

import (
	"testing"
)

func TestOrangeRotting(t *testing.T) {
	got := orangesRotting([][]int{{2, 2}, {1, 1}, {0, 0}, {2, 0}})
	want := 1
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
