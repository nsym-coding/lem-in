package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	path     []*Room
	visited  bool
	occupied bool
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

func NumAnts(s []string) int {
	antNum := s[0]
	intAntNum, _ := strconv.Atoi(antNum)
	s = readAntsFile("ants.txt")
	if s[0] <= "0" {
		err := fmt.Errorf("invalid number of ants")
		fmt.Println(err.Error())
	}
	return intAntNum
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

var (
	StartR = StartRoom(readAntsFile("ants.txt"))
	EndR   = EndRoom(readAntsFile("ants.txt"))
)

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

// func doesContain(s string, sl []string) bool {
// 	for _, word := range sl {
// 		if s == word {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {

	test := Graph{}

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

	ants := Ants{}
	test.Print()
	DFS(test.getRoom(StartR), test)
	ants.Output()
}

func (g *Graph) PrintPath() {
	fmt.Println(StartRoom(readAntsFile("ants.txt")))
	for _, v := range g.rooms {
		for _, r := range v.path {
			fmt.Println(r)
		}
	}
}

// add all edges
// Add edge to the graph. deals with a directional graph only
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
	} else if fromRoom.key == EndR {
		toRoom.adjacent = append(toRoom.adjacent, fromRoom)
	} else if toRoom.key == StartR {
		toRoom.adjacent = append(toRoom.adjacent, fromRoom)
	} else {
		fromRoom.adjacent = append(fromRoom.adjacent, toRoom)
	}
	//add edge etc
}

//Print will print the adjacent list for each Room of the graph
func (g *Graph) Print() {
	// fmt.Println(readAntsFile("ants.txt"))
	fmt.Printf("The number of ants is: %v ", NumAnts(readAntsFile("ants.txt")))
	fmt.Println()

	for _, v := range g.rooms {
		if v.key == StartR {
			fmt.Printf("\n Start Room is %v : ", StartR)

		} else if v.key == EndR {
			fmt.Printf("\n End Room is %v :", EndR)

		} else {
			fmt.Printf("\n Room %v : ", v.key)

		}
		for _, v := range v.adjacent {
			fmt.Printf(" %v,", v.key)
		}
	}
	fmt.Println()
}

// global variable which will store all of the valid paths in a slice of slices of string.
var validPaths [][]*Room

// Depth first search function that operates recursively
func DFS(r *Room, g Graph) {
	// vList := []string{}
	sRoom := g.getRoom(StartR)

	// set the room being checked visited status to true
	// range through the neighbours of the r
	if r.key != EndR {
		r.visited = true
		for _, nbr := range r.adjacent {
			if !nbr.visited {
				/* for each neighbour that hasn't been visited,
				- append their key to the visited slice,
				- then apply dfs to them recursively,
				- then append their key to their path value
				*/

				nbr.path = append(r.path, nbr)
				if contains(nbr.path, EndR) {
					validPaths = append(validPaths, nbr.path)
				}
				DFS(nbr, Graph{g.rooms})
			}
		}
	} else if !contains(sRoom.adjacent, EndR) {
		sRoom.adjacent = sRoom.adjacent[1:]
		DFS(sRoom, Graph{g.rooms})
	}

}

type Ants struct {
	antz []*Ant
}
type Ant struct {
	key  string
	path []*Room // valid path
	// currentRoom Room
}

func (a *Ants) Output() {
	numOfAnts := NumAnts(readAntsFile("ants.txt"))
	// valid paths from dfs function
	unmovedAnts := []string{}

	for i := 1; i <= numOfAnts; i++ {
		unmovedAnts = append(unmovedAnts, strconv.Itoa(i))
		a.antz = append(a.antz, &Ant{key: "L" + strconv.Itoa(i)})
		a.antz = append(a.antz, &Ant{path: validPaths[0]})
	}

	for _, str := range a.antz {
		for _, room := range str.path {
			str.path[0].occupied = true
			if !room.occupied {
				fmt.Println(room.key)
			}
		}
	}

	// fmt.Println(unmovedAnts)
	// vp := validPaths

	// a := Ant{}

	// for i := range unmovedAnts {
	// 	a.key = "L" + unmovedAnts[i]
	// 	a.path = validPaths[0]
	// }
	// need to use unmoved and movingants

	// if room is unoccupied, add ant into moving slice
	// if all rooms upto that point are occupied, turn done, println
	// for range unmovedAnts {
	// 	for _, path := range vp {
	// 		for v, room := range path {
	// 			movingAnts = append(movingAnts, unmovedAnts[0])
	// 			unmovedAnts = unmovedAnts[1:]
	// 			for i, ant := range movingAnts {
	// 				if room == EndR {
	// 					movingAnts = movingAnts[1:]
	// 					// fmt.Print(movingAnts)
	// 				}
	// 				fmt.Printf("L%v-%v ", ant, room)
	// 				if room != path[0] && i == len(movingAnts)-1 {
	// 					// fmt.Println(movingAnts)
	// 					room = path[v-1]
	// 				}
	// 			}
	// 			fmt.Println()
	// 		}
	// 	}
	// }
}
