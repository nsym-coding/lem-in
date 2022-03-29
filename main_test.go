package main

import (
	"strconv"
	"testing"
)

// func Test_readAntsFile(t *testing.T) {
// 	got := readAntsFile("ants.txt")[0]
// 	want := "10"

// 	if got != want {
// 		t.Errorf("got: %q, wanted: %q", got, want)
// 	}
// }

func TestNumAnts(t *testing.T) {

	got := NumAnts(readAntsFile("example00.txt"))
	want, _ := strconv.Atoi("4")

	if got != want {
		t.Errorf("got: %q, wanted: %q instead", got, want)
	}
}

func TestStartRoom(t *testing.T) {

	got := StartRoom(readAntsFile("example00.txt"))
	want := "0"
	if got != want {
		t.Errorf("got: %q, wanted: %q", got, want)
	}
}

// func TestEndRoom(t *testing.T) {
// 	got := EndRoom(readAntsFile("ants.txt"))
// 	want := "end"

// 	if got != want {
// 		t.Errorf("got %q, wanted: %q", got, want)
// 	}
// }
