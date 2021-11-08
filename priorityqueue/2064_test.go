package priorityqueue

import (
	"reflect"
	"testing"
)

func TestMinimizedMaximum(t *testing.T) {
	got := minimizedMaximum(6, []int{11, 6})
	want := 3
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinimizedMaximum2(t *testing.T) {
	got := minimizedMaximum(7, []int{15, 10, 10})
	want := 5
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinimizedMaximum3(t *testing.T) {
	got := minimizedMaximum(1, []int{100000})
	want := 100000
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinimizedMaximum4(t *testing.T) {
	got := minimizedMaximum(7, []int{9, 1, 8, 12})
	want := 6
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMinimizedMaximum5(t *testing.T) {
	got := minimizedMaximum(7, []int{9, 1, 5, 12})
	want := 5
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// 5 4
// 1
// 4 4
// 6 6
