package main

import "testing"

func Test_NumAnts(t *testing.T) {

	got := NumAnts(readAntsFile("example01.txt"))
	want := 10

	if got != want {
		t.Errorf("got: %v, wanted: %v instead", got, want)
	}
}

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
