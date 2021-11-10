package dynamic

import (
	"reflect"
	"testing"
)

func TestPossiblyEquals(t *testing.T) {
	got := possiblyEquals("internationalization", "i18n")
	want := true
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPossiblyEquals2(t *testing.T) {
	got := possiblyEquals("l123e", "44")
	want := true
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPossiblyEquals3(t *testing.T) {
	got := possiblyEquals("a5b", "c5b")
	want := false
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPossiblyEquals4(t *testing.T) {
	got := possiblyEquals("112s", "g841")
	want := true
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPossiblyEquals5(t *testing.T) {
	got := possiblyEquals("ab", "a2")
	want := false
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPossiblyEquals6(t *testing.T) {
	got := possiblyEquals("ab", "a2")
	want := false
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
