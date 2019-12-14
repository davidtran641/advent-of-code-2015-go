package main
import (
	"testing"
)

func TestCountHouse(t *testing.T) {
	assertCountHouse(t, ">", 2)
	assertCountHouse(t, "^>v<", 4)
	assertCountHouse(t, "^v^v^v^v^v", 2)
}

func assertCountHouse(t *testing.T, in string, want int) {
	out := CountHouse(in)
	if out != want {
		t.Errorf("CountHouse(%v) = %v, want %v", in, out, want)
	}
}

func TestCountHouseReal(t *testing.T) {
	in := ReadFile("input.txt")
	out := CountHouse(in)
	want := 2572

	if out != want {
		t.Errorf("CountHouse = %v, want %v", out, want)
	}
}

func TestCountHouseRobo(t *testing.T) {
	assertCountHouseRobo(t, "^v", 3)
	assertCountHouseRobo(t, "^>v<", 3)
	assertCountHouseRobo(t, "^v^v^v^v^v", 11)
}

func assertCountHouseRobo(t *testing.T, in string, want int) {
	out := CountHouseRobo(in)
	if out != want {
		t.Errorf("CountHouseRobo(%v) = %v, want %v", in, out, want)
	}
}

func TestCountHouseRoboReal(t *testing.T) {
	in := ReadFile("input.txt")
	out := CountHouseRobo(in)
	want := 0

	if out != want {
		t.Errorf("CountHouseRobo = %v, want %v", out, want)
	}
}