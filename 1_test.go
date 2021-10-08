package main

import (
	"reflect"
	"testing"
)

func TestTwoSumDifferentNums(t *testing.T) {
	nums := []int{10, 2, 3, 4, 5, 6}
	got := twoSum(nums, 5)
	want := []int{1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTwoSumEqualNums(t *testing.T) {
	nums := []int{10, 4, 2, 30, 4, 5, 16}
	got := twoSum(nums, 8)
	want := []int{1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
