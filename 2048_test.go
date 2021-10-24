package main

import (
	"reflect"
	"testing"
)

func TestNextBeautifulNumber(t *testing.T) {
	got := nextBeautifulNumber(1)
	want := 22
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNextBeautifulNumber2(t *testing.T) {
	got := nextBeautifulNumber(22)
	want := 122
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestNextBeautifulNumber3(t *testing.T) {
	got := nextBeautifulNumber(666665)
	want := 666666
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
