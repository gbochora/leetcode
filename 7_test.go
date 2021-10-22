package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	got := reverse(123)
	want := 321
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
