package linkedlist

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{ArrayToListInt([]int{1, 4, 5}), ArrayToListInt([]int{1, 3, 4}), ArrayToListInt([]int{2, 6})}

	got := mergeKLists(lists)
	want := ArrayToListInt([]int{1, 1, 2, 3, 4, 4, 5, 6})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMergeKLists2(t *testing.T) {
	lists := []*ListNode{nil, ArrayToListInt([]int{1, 4, 5}), ArrayToListInt([]int{0, 0, 1, 3, 4}), ArrayToListInt([]int{2, 6}), nil}

	got := mergeKLists(lists)
	want := ArrayToListInt([]int{0, 0, 1, 1, 2, 3, 4, 4, 5, 6})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMergeKListsNilList(t *testing.T) {
	lists := []*ListNode{nil}

	got := mergeKLists(lists)
	var want *ListNode
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMergeKListsEmptyList(t *testing.T) {
	lists := []*ListNode{}

	got := mergeKLists(lists)
	var want *ListNode
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
