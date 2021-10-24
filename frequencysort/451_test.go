package frequencysort

import (
	"reflect"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	got := frequencySort("treerr")
	want := "rrreet"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
