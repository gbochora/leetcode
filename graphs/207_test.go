package graphs

import (
	"reflect"
	"testing"
)

func TestCanFinish(t *testing.T) {
	got := canFinish(2, [][]int{{1, 0}, {0, 1}})
	want := false
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}
