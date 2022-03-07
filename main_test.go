package main

import (
	"testing"
)

func Test_readAntsFile(t *testing.T) {
	got := readAntsFile("ants.txt")[0]
	want := "10"

	if got != want {
		t.Errorf("got: %q, wanted: %q", got, want)
	}
}

func TestNumAnts(t *testing.T) {

	got := NumAnts(readAntsFile("ants.txt"))
	want := 10

	if got != want {
		t.Errorf("got: %q, wanted: %q instead", got, want)
	}
}

func TestStartRoom(t *testing.T) {

	got := StartRoom(readAntsFile("ants.txt"))
	want := "start"
	if got != want {
		t.Errorf("got: %q, wanted: %q", got, want)
	}
}

func TestEndRoom(t *testing.T) {
	got := EndRoom(readAntsFile("ants.txt"))
	want := "end"

	if got != want {
		t.Errorf("got %q, wanted: %q", got, want)
	}
}
