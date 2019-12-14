package main
import (
	"testing"
)

func TestCountBracket(t *testing.T) {
	assertCountBracket(t, "(())", 0)
	assertCountBracket(t, "(())", 0)
	assertCountBracket(t, "(()(()(", 3)
}

func TestCountBracketReal(t *testing.T) {
	in := ReadFile("input.txt")
	assertCountBracket(t, in, 138)
}

func assertCountBracket(t *testing.T, in string, want int) {
	out := CountOpenBracket(in)
	if out != want {
		t.Errorf("CountOpenBracket(%v) = %v, want %v", in, out, want)
	}
}

func TestFindBasementPosition(t *testing.T) {
	assertFindBasementPosition(t, ")", 1)
	assertFindBasementPosition(t, "()())", 5)
}

func TestFindBasementPositionReal(t *testing.T) {
	in := ReadFile("input.txt")
	assertFindBasementPosition(t, in, 1771)
}

func assertFindBasementPosition(t *testing.T, in string, want int) {
	out := FindBasementPosition(in)
	if out != want {
		t.Errorf("FindBasementPosition(%v) = %v, want %v", in, out, want)
	}
}