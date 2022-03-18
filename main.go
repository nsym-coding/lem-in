package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Graph structure
type Graph struct {
	rooms []*Room
}

// Room structure
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
	s = readAntsFile("ants.txt")
	antNum := s[0]
	if s[0] <= "0" {
		err := fmt.Errorf("invalid number of ants")
		fmt.Println(err.Error())
	}
	n, _ := strconv.Atoi(antNum)

	return n
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
	// fmt.Println(startRoom)
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
	// fmt.Println(endRoom)
	return endRoom
}

var (
	StartR = StartRoom(readAntsFile("ants.txt"))
	EndR   = EndRoom(readAntsFile("ants.txt"))
)

var pathSlice [][]*Room

var start *Room

// Add Room to a graph
func (g *Graph) AddRoom(k string) {
	if contains(g.rooms, k) {
		err := fmt.Errorf("Room %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.rooms = append(g.rooms, &Room{key: k})
	}
}

// getRoom returns a pointer to the Room key integer
func (g *Graph) getRoom(k string) *Room {
	for i, v := range g.rooms {
		if v.key == k {
			return g.rooms[i]
		}
	}
	return nil
}

// contains checks if the Room key exists
func contains(s []*Room, k string) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func doesContain(s string, sl []string) bool {
	for _, word := range sl {
		if s == word {
			return true
		}
	}
	return false
}

// delete edge from starting room
// func (r *Room) DeleteEdge(k string) {

// 	// if !contains(g.rooms, k) {
// 	//     err := fmt.Errorf("edge %v not deleted because it doesn't exist", k)
// 	//     fmt.Println(err.Error())
// 	// } else {
// 	fmt.Println("Error check 1")

// 	// for _, t := range r{
// 	// 	fmt.Println("Error check 2")

// 	start := g.get

// 	for r.key = k{
// 		r.adjacent =
// 	}
// 	fmt.Println("Error check 3")

// 	r.adjacent = r.adjacent[1:]
// 	fmt.Println("something")
// 	// for i , room := range g.rooms[i].adjacent{
// 	//     if  room.key == k {
// 	//         g.rooms[i].adjacent =
// 	//     }
// 	// }

// }

func main() {
	// err := errors.New("ERROR: invalid data format")
	// if err != nil {
	// 	fmt.Print(err, "\n")
	// 	os.Exit(1)
	// }

	test := Graph{}

	// adding all rooms
	for i, line := range readAntsFile("ants.txt") {
		if strings.Contains(string(line), " ") {
			test.AddRoom(strings.Split(readAntsFile("ants.txt")[i], " ")[0])
		}
		// adding all edges from and to rooms
		// maybe add a condition so that it adds the edges in order i.e. the end room as the last edge?
		if strings.Contains(string(line), "-") {
			test.AddEdge(strings.Split(readAntsFile("ants.txt")[i], "-")[0], strings.Split(readAntsFile("ants.txt")[i], "-")[1])
			// test.AddEdge(strings.Split(readAntsFile("ants.txt")[i], "-")[1], strings.Split(readAntsFile("ants.txt")[i], "-")[0])
		}

	}

	test.Print()
	DFS(test.getRoom(StartR), test)
	output(pathSlice)
	// dfsStart(test)
	// GFS(test)
	// test.PrintPath()
	// DeleteEdge()
	// test.Print()
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
	// get Room
	fromRoom := g.getRoom(from)
	toRoom := g.getRoom(to)

	// check error

	// if toRoom.key == "E2" && fromRoom.key != "B1" {
	// 	toRoom.adjacent = append(toRoom.adjacent, fromRoom)
	// }
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
	// add edge etc
}

// Print will print the adjacent list for each Room of the graph
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

// Depth first search function that operates recursively
func DFS(r *Room, g Graph) {
	vList := []string{}
	sRoom := g.getRoom(StartR)

	// set the room being checked visited status to true
	if r.key != EndR {
		r.visited = true

		// append the r key to the visited list
		vList = append(vList, r.key)

		// range through the neighbours of the r
		for _, nbr := range r.adjacent {
			if !nbr.visited {
				/* for each neighbour that hasn't been visited,
				- append their key to the visited slice,
				- then apply dfs to them recursively,
				- then append their key to their path value
				*/

				// fmt.Println("*", vList)

				nbr.path = append(r.path, nbr)
				if contains(nbr.path, EndR) {
					// for _, r := range nbr.path {
					// 	fmt.Println(r.key)
					// }
					pathSlice = append(pathSlice, nbr.path)
				}
				// fmt.Println(nbr.path)
				vList = append(vList, nbr.key)
				DFS(nbr, Graph{g.rooms})

			}
		}

	} else {
		if len(sRoom.adjacent) > 1 && !contains(sRoom.adjacent, EndR) {
			vList = append(vList, r.key)

			// fmt.Println("*", vList)
			sRoom.adjacent = sRoom.adjacent[1:][:]

			// for _, v := range sRoom.adjacent {
			// 	fmt.Println(v.key)
			// }
			DFS(sRoom, Graph{g.rooms})

		} else {
			vList = append(vList, r.key)
			// fmt.Println("*", vList)
		}
	}
}

func output(pathSlice [][]*Room) {
	// fmt.Println(len(pathSlice))
	ants := NumAnts(readAntsFile("ants.txt"))

	var unmovedAnts []string

	for i := 1; i <= ants; i++ {
		unmovedAnts = append(unmovedAnts, strconv.Itoa(i))
	}
	// fmt.Println(unmovedAnts)
	// map to hold each visited rooms
	// occupied := make(map[string]bool)

	// for _, path := range pathSlice {
	// 	for _, node := range path {
	// 		fmt.Print(node.key)
	// 	}
	// }

	// if room is unoccupied, add ant into movingAnt slice
	// if all rooms up to that point are occupied, turn done, println
	var movingAnts []string

	for range unmovedAnts {
		for _, path := range pathSlice {
			for i, node := range path {
				if !node.occupied {
					movingAnts = append(movingAnts, unmovedAnts[0])
					unmovedAnts = unmovedAnts[1:]
					// fmt.Println(unmovedAnts)
					// fmt.Println(movingAnts)
					fmt.Printf("%v room is unoccupied\n",node.key)
					for _, v := range movingAnts {
						fmt.Printf("L%v-%v ", v, node.key)
						//node.occupied = true
						// fmt.Print(v)
						if node.key == EndR {
							movingAnts = movingAnts[1:]
							// fmt.Print(movingAnts)
						}
						if node != path[0] {
							node = path[i-1]
							// fmt.Println(node.path)
						}
					}
					node.occupied = true
					// turn
					fmt.Println()
					// node.occupied = false
					} else {
					for _, v := range movingAnts {
						fmt.Printf("%v room is occupied by ant L%v ", node.key, v)
					}
					fmt.Println()
				}
			}
		}
	}
}

// Function that initialises that DFS algorithm by taking the target graph as an argument
// func dfsStart(g *Graph) {
// 	DFS(g.getRoom(StartR))
// }
