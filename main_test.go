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

func Test_output(t *testing.T) {
	type args struct {
		pathSlice [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output(tt.args.pathSlice)
		})
	}
}
