package threesum

import (
	"reflect"
	"testing"
)

func TestCountValidWords2(t *testing.T) {
	got := threeSum([]int{-1, 0, 1, 2, -1, -4})
	want := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
