package main

import (
	"reflect"
	"testing"
)

func TestGraph_Print(t *testing.T) {
	type fields struct {
		rooms []*Room
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				rooms: tt.fields.rooms,
			}
			g.Print()
		})
	}
}

func Test_readAntsFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readAntsFile(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readAntsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumAnts(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumAnts(tt.args.s); got != tt.want {
				t.Errorf("NumAnts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_getRoom(t *testing.T) {
	type fields struct {
		rooms []*Room
	}
	type args struct {
		k string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Room
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				rooms: tt.fields.rooms,
			}
			if got := g.getRoom(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Graph.getRoom() = %v, want %v", got, tt.want)
			}
		})
	}
}
