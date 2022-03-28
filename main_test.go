package main

import (
	"testing"
)

func Test_Increment(t *testing.T) {

	a := 1
	b := 2
	var c []int
	c = append(c, a)
	c = append(c, b)
	var d [][]int
	d = append(d, c)

	got := Increment(d, 1)[0][0]
	want := 2

	if got != want {
		t.Errorf("got: %v, wanted: %v instead", got, want)
	}
}

func Test_readAntsFile(t *testing.T) {
	got := readAntsFile("example01.txt")[0]
	want := "10"

	if got != want {
		t.Errorf("got: %q, wanted: %q", got, want)
	}
}

func TestNumAnts(t *testing.T) {

	tests := []struct {
		name string
		s    []string
		want int
	}{
		{name: "example00", s: readAntsFile("example01.txt"), want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumAnts(tt.s); got != tt.want {
				t.Errorf("NumAnts() = %v, want %v", got, tt.want)
			}
		})
	}
}
