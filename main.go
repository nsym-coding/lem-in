package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Graph structure
type Graph struct {
	rooms []*Room
}

//Room structure
type Room struct {
	key      string
	adjacent []*Room
}

// Reads file and returns a string slice
func readAntsFile(filename string) []string {
	file, _ := os.Open(filename)
	scanned := bufio.NewScanner(file)
	scanned.Split(bufio.ScanLines)

	var lines []string

	for scanned.Scan() {
		lines = append(lines, scanned.Text())
	}
	return lines
}

func NumAnts(s []string) string {
	antNum := s[0]
	s = readAntsFile("ants.txt")
	if s[0] <= "0" {
		err := fmt.Errorf("invalid number of ants")
		fmt.Println(err.Error())
	}
	return antNum
}

// Gets out the start room and returns it
func StartRoom([]string) string {

	var startRoom string
	s := readAntsFile("ants.txt")
	//	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		if s[i] == "##start" {
			startRoom = strings.Split(string(s[i+1]), " ")[0]
		}

	}
	//fmt.Println(startRoom)

	return startRoom

}

// Gets out the end room and returns it
func EndRoom([]string) string {
	var endRoom string
	s := readAntsFile("ants.txt")
	// fmt.Println(s)

	for i := 0; i < len(s); i++ {
		if s[i] == "##end" {
			endRoom = strings.Split(string(s[i+1]), " ")[0]
		}

	}
	//fmt.Println(endRoom)

	return endRoom
}

//Add Room to a graph
func (g *Graph) AddRoom(k string) {
	if contains(g.rooms, k) {
		err := fmt.Errorf("Room %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.rooms = append(g.rooms, &Room{key: k})
	}
}

//getRoom returns a pointer to the Room key integer
func (g *Graph) getRoom(k string) *Room {
	for i, v := range g.rooms {
		if v.key == k {
			return g.rooms[i]
		}
	}
	return nil
}

//contains checks if the Room key exists

func contains(s []*Room, k string) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func main() {
	StartRoom(readAntsFile("ants.txt"))
	EndRoom(readAntsFile("ants.txt"))

	test := &Graph{}
	// rooms := []string{}
	// for _, room := range rooms {
	// 	test.AddRoom(room)
	// }

	//fmt.Println(readAntsFile("ants.txt"))

	//add rooms
	// fmt.Println(strings.Split(readAntsFile("ants.txt")[2], " ")[0])

	// reading all rooms

	/* need to add in condition so that it only selects the room after ##start as the start room
	and the room after ##end as the end room*/
	//var startRoom string
	//var endRoom string
	// for _, line := range readAntsFile("ants.txt") {
	// 	if strings.Contains(string(line), "##start") {
	// 		line = startRoom
	// 		fmt.Println("This is the start room: %s", startRoom)

	// 	}

	//adding all rooms
	for i, line := range readAntsFile("ants.txt") {
		if strings.Contains(string(line), " ") {
			test.AddRoom(strings.Split(readAntsFile("ants.txt")[i], " ")[0])
		}
		// adding all edges from and to rooms
		// maybe add a condition so that it adds the edges in order i.e. the end room as the last edge?
		if strings.Contains(string(line), "-") {
			test.AddEdge(strings.Split(readAntsFile("ants.txt")[i], "-")[0], strings.Split(readAntsFile("ants.txt")[i], "-")[1])
			//test.AddEdge(strings.Split(readAntsFile("ants.txt")[i], "-")[1], strings.Split(readAntsFile("ants.txt")[i], "-")[0])

		}

	}
	test.Print()
}

// add all edges

//Add edge to the graph. deals with a directional graph only

func (g *Graph) AddEdge(from, to string) {
	//get Room
	fromRoom := g.getRoom(from)
	toRoom := g.getRoom(to)

	//check error
	if fromRoom == nil || toRoom == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromRoom.adjacent, to) {
		err := fmt.Errorf("existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if fromRoom == toRoom {
		err := fmt.Errorf("cannot connect room to itself (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else if fromRoom.key == EndRoom(readAntsFile("ants.txt")) {
		toRoom.adjacent = append(toRoom.adjacent, fromRoom)
	} else if toRoom.key == StartRoom(readAntsFile("ants.txt")) {
		toRoom.adjacent = append(toRoom.adjacent, fromRoom)
	} else {
		fromRoom.adjacent = append(fromRoom.adjacent, toRoom)
	}
	//add edge
}

//Print will print the adjacent list for each Room of the graph

func (g *Graph) Print() {
	// fmt.Println(readAntsFile("ants.txt"))

	fmt.Printf("The number of ants is: %v ", NumAnts(readAntsFile("ants.txt")))
	fmt.Println()

	for _, v := range g.rooms {
		if v.key == StartRoom(readAntsFile("ants.txt")) {
			fmt.Printf("\n Start Room is %v : ", StartRoom(readAntsFile("ants.txt")))

		} else if v.key == EndRoom(readAntsFile("ants.txt")) {
			fmt.Printf("\n End Room is %v :", EndRoom(readAntsFile("ants.txt")))

		} else {
			fmt.Printf("\n Room %v : ", v.key)

		}
		for _, v := range v.adjacent {
			fmt.Printf(" %v", v.key)
		}
	}
	fmt.Println()
}

func (g *Graph) dfs(*Graph) {

}
