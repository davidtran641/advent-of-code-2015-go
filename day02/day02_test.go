package main
import (
	"testing"
)

func TestCountBracket(t *testing.T) {
	assertFindWrappingArea(t, "2x3x4", 58)
	assertFindWrappingArea(t, "1x1x10", 43)
}

func TestCountBracketReal(t *testing.T) {
	in := ReadLines("input.txt")
	out := FindWrappingAreaList(in)
	want := 1586300

	if out != want {
		t.Errorf("FindWrappingAreaList = %v, want %v", out, want)
	}
}

func assertFindWrappingArea(t *testing.T, in string, want int) {
	out := FindWrappingArea(in)
	if out != want {
		t.Errorf("FindWrappingArea(%v) = %v, want %v", in, out, want)
	}
}

func TestFindRobbinLength(t *testing.T) {
	assertFindRobbinLength(t, "2x3x4", 34)
	assertFindRobbinLength(t, "1x1x10", 14)
}

func assertFindRobbinLength(t *testing.T, in string, want int) {
	out := FindRobbinLength(in)
	if out != want {
		t.Errorf("FindRobbinLength(%v) = %v, want %v", in, out, want)
	}
}

func TestFindRobbinLengthList(t *testing.T) {
	in := ReadLines("input.txt")
	out := FindRobbinLengthList(in)
	want := 3737498

	if out != want {
		t.Errorf("FindWrappingAreaList = %v, want %v", out, want)
	}
}