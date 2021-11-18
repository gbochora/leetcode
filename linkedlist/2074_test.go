package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverseEvenGroup(t *testing.T) {
	got := reverseEvenLengthGroups(ArrayToListInt([]int{5, 2, 6, 3, 9, 1, 7, 3, 8, 4}))
	want := ArrayToListInt([]int{5, 6, 2, 3, 9, 1, 4, 8, 3, 7})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseEvenGroup1(t *testing.T) {
	got := reverseEvenLengthGroups(ArrayToListInt([]int{1, 2}))
	want := ArrayToListInt([]int{1, 2})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseEvenGroup3(t *testing.T) {
	got := reverseEvenLengthGroups(ArrayToListInt([]int{1, 2, 3, 4, 5}))
	want := ArrayToListInt([]int{1, 3, 2, 5, 4})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseEvenGroup4(t *testing.T) {
	got := reverseEvenLengthGroups(ArrayToListInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	want := ArrayToListInt([]int{1, 3, 2, 4, 5, 6, 7, 8, 9})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseEvenGroup5(t *testing.T) {
	got := reverseEvenLengthGroups(ArrayToListInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	want := ArrayToListInt([]int{1, 3, 2, 4, 5, 6, 10, 9, 8, 7})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseEvenGroup6(t *testing.T) {
	got := reverseEvenLengthGroups(ArrayToListInt([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}))
	want := ArrayToListInt([]int{1, 3, 2, 4, 5, 6, 10, 9, 8, 7, 12, 11})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseEvenGroup7(t *testing.T) {
	got := reverseEvenLengthGroups(ArrayToListInt([]int{1, 2, 3, 4}))
	want := ArrayToListInt([]int{1, 3, 2, 4})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseEvenGroup8(t *testing.T) {
	got := reverseEvenLengthGroups(ArrayToListInt([]int{1}))
	want := ArrayToListInt([]int{1})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseEvenGroup9(t *testing.T) {
	got := reverseEvenLengthGroups(ArrayToListInt([]int{4, 3, 0, 5, 1, 2, 7, 8, 6}))
	want := ArrayToListInt([]int{4, 0, 3, 5, 1, 2, 7, 8, 6})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
