package stringsmanipulation

import (
	"reflect"
	"testing"
)

func TestLongestSubstring(t *testing.T) {
	got := lengthOfLongestSubstring("abcabcbb")
	want := 3
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLongestSubstring2(t *testing.T) {
	got := lengthOfLongestSubstring("aaaa12345234098")
	want := 7
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
