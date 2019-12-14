package main
import (
	"testing"
)

func TestFindSuitableHash(t *testing.T) {
	assertFindSuitableHash(t, "abcdef", 5, 609043)
	assertFindSuitableHash(t, "pqrstuv", 5, 1048970)
	assertFindSuitableHash(t, "yzbqklnj", 5, 282749)
	assertFindSuitableHash(t, "yzbqklnj", 6, 9962624)
}

func assertFindSuitableHash(t *testing.T, in string, numZero int, want int) {
	out := FindSuitableHash(in, numZero)
	if out != want {
		t.Errorf("FindSuitableHash(%v) = %v, want %v", in, out, want)
	}
}
