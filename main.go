package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Ants struct {
	antz []*Ant
}

type Ant struct {
	key         string
	path        []*Room
	currentRoom Room
}

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

	args := os.Args[1]


	antNum := s[0]
	s = readAntsFile(args)
	if s[0] <= "0" {
		err := fmt.Errorf("ERROR: Invalid number of ants")
		fmt.Println(err.Error())
	}

	antNumInt, _ := strconv.Atoi(antNum)
	return antNumInt
}

// Gets out the start room and returns it
func StartRoom([]string) string {

	args := os.Args[1]

	var startRoom string
	s := readAntsFile(args)

	for i := 0; i < len(s); i++ {
		if s[i] == "##start" {
			startRoom = strings.Split(string(s[i+1]), " ")[0]
		}

	}

	return startRoom

}

// Gets out the end room and returns it
func EndRoom([]string) string {

	args := os.Args[1]

	var endRoom string
	s := readAntsFile(args)

	for i := 0; i < len(s); i++ {
		if s[i] == "##end" {
			endRoom = strings.Split(string(s[i+1]), " ")[0]
		}

	}
	return endRoom
}

var (
	args = os.Args[1]
	StartR = StartRoom(readAntsFile(args))
	EndR   = EndRoom(readAntsFile(args))
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

func doesContain(s string, sl []string) bool {
	for _, word := range sl {
		if s == word {
			return true
		}
	}
	return false
}

func doesContainRoom(sl []*Room, s string) bool {

	for _, word := range sl {
		if s == word.key {
			return true
		}
	}
	return false
}

func (g *Graph) PrintPath() {
	fmt.Println(StartRoom(readAntsFile(args)))
	for _, v := range g.rooms {
		for _, r := range v.path {
			fmt.Println(r)
		}
	}
}

// add all edges
// Add edge to the graph. deals with a directional graph only but condition in the main makes it undirected
func (g *Graph) AddEdge(from, to string) {
	//get Room
	fromRoom := g.getRoom(from)
	toRoom := g.getRoom(to)

	//check error
	if fromRoom == nil || toRoom == nil {
		err := fmt.Errorf("ERROR: invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
		os.Exit(0)
	} else if contains(fromRoom.adjacent, to) {
		err := fmt.Errorf("ERROR: existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
		os.Exit(0)
	} else if fromRoom == toRoom {
		err := fmt.Errorf("ERROR: cannot connect room to itself (%v --> %v)", from, to)
		fmt.Println(err.Error())
		os.Exit(0)
	} else if fromRoom.key == EndR {
		//toRoom.adjacent = append(toRoom.adjacent, fromRoom)
		//} //else if toRoom.key == StartR {
		//toRoom.adjacent = append(toRoom.adjacent, fromRoom)
	} else {
		fromRoom.adjacent = append(fromRoom.adjacent, toRoom)
	}
	//add edge etc
}

//Print will print the adjacent list for each Room of the graph
func (g *Graph) Print() {
	// fmt.Println(readAntsFile(args))
	fmt.Printf("The number of ants is: %v ", NumAnts(readAntsFile(args)))
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

// Slices to hold paths from both algos for comparison
var dfsPaths [][]*Room
var bfsPaths [][]*Room

// Depth First Search algorithm that operates recursively
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

				nbr.path = append(r.path, nbr)
				if contains(nbr.path, EndR) {

					dfsPaths = append(dfsPaths, nbr.path)

				}

				vList = append(vList, nbr.key)

				DFS(nbr, Graph{g.rooms})

			}

		}

	} else {

		if len(sRoom.adjacent) > 1 && !contains(sRoom.adjacent, EndR) {
			vList = append(vList, r.key)

			sRoom.adjacent = sRoom.adjacent[1:][:]

			DFS(sRoom, Graph{g.rooms})

		} else {
			vList = append(vList, r.key)

		}
	}
	dfsPaths = PathDupeCheck(dfsPaths)

}

// Depth first search function that operates recursively
func DFSBFS(r *Room, g Graph) bool {

	//vList := []string{}
	//sRoom := g.getRoom(StartR)

	// set the room being checked visited status to true
	if r.key != EndR {
		r.visited = true

		// range through the neighbours of the r
		for _, nbr := range r.adjacent {
			if !nbr.visited {
				/* for each neighbour that hasn't been visited,
				- append their key to the visited slice,
				- then apply dfs to them recursively,
				- then append their key to their path value
				*/

				nbr.path = append(r.path, nbr)
				if contains(nbr.path, EndR) {

					return true

				}

			}

		}

	}
	return false
}

// Function to remove room using its key
func RemoveRoomIndex(s []*Room, index string) []*Room {
	ret := make([]*Room, 0)

	for i := 0; i < len(s); i++ {
		if s[i].key == index {
			ret = append(ret, s[:i]...)
			ret = append(ret, s[i+1:]...)
		}
	}
	return ret
}

// function to remove an element from a slice of ints
func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

//removes a string from a slice (unordered)
func remove(s []*Room, k string) []*Room {
	for i := 0; i < len(s); i++ {
		if s[i].key == k {
			s[i] = s[len(s)-1]

		}

	}
	return s[:len(s)-1]
}

//delete edge from starting room
func DeleteEdge(r *Room, g Graph) {

	for i := 0; i < len(r.path); i++ {
		for _, room := range g.rooms {
			//	for _ , edge := range room.adjacent
			for j := 0; j < len(room.adjacent); j++ {
				if room.adjacent[j] == r.path[i] {
					room.adjacent = remove(room.adjacent, r.key)
				}
			}
		}
	}
}

// Breadth First Search algorithm
func BFS(r *Room, g Graph) {

	var vPaths [][]*Room

	//queue variable, procedurally populated with rooms yet to be visited
	var queue []*Room

	//set start room as visited
	r.visited = true

	//initialise queue with start room
	queue = append(queue, r)

	// }

	// checks if there is a link between start and end directly
	for i, v := range g.getRoom(StartR).adjacent {
		if v.key == g.getRoom(EndR).key {
			g.getRoom(EndR).path = append(g.getRoom(EndR).path, g.getRoom(StartR))
			vPaths = append(vPaths, g.getRoom(StartR).path)
			g.getRoom(StartR).adjacent = append(g.getRoom(StartR).adjacent[:i], g.getRoom(StartR).adjacent[i+1:]...)
		}

	}

	//checks the queue for the end room and if the queue is not empty

	for !contains(queue, g.getRoom(EndR).key) && len(queue) >= 1 {
		qfront := queue[0]

		for _, room := range qfront.adjacent {
			if !room.visited {
				room.visited = true
				room.path = append(qfront.path, room)
				//
				queue = append(queue, room)
			}

		}

		queue = queue[1:]

		if doesContainRoom(queue, g.getRoom(EndR).key) {

			for _, room := range g.rooms {
				room.visited = false
			}
			vPaths = append(vPaths, qfront.path)

			for _, r := range qfront.path {
				DeleteEdge(r, g)

			}
			if len(g.getRoom(StartR).adjacent) == 0 {

				break
			}

			if len(g.getRoom(StartR).adjacent) >= 1 {
				for _, froom := range g.getRoom(StartR).adjacent {
					for _, sroom := range froom.adjacent {
						if sroom.key != g.getRoom(EndR).key {
							break
						} else {
							BFS(g.getRoom(StartR), Graph{g.rooms})
							queue = queue[1:]
						}
					}
				}
			}
			BFS(g.getRoom(StartR), Graph{g.rooms})

		}
	}
	for _, v := range vPaths {
		v = append(v, g.getRoom(EndR))
		bfsPaths = append(bfsPaths, v)
	}
	bfsPaths = PathDupeCheck(bfsPaths)

}

//returns the optimal path between bfs & dfs algos
func PathSelection(bfs [][]*Room, dfs [][]*Room) [][]*Room {

	bfsPathNum := len(bfs)
	dfsPathNum := len(dfs)

	if bfsPathNum > dfsPathNum {
		validPaths = append(validPaths, bfsPaths...)
	} else if dfsPathNum > bfsPathNum {
		validPaths = PathDupeCheck(append(validPaths, dfsPaths...))
	} else {

		bfscounter := 0

		dfscounter := 0

		for _, path := range bfs {

			bfscounter += len(path)

		}

		for _, path := range dfs {
			dfscounter += len(path)
		}

		if bfscounter < dfscounter {
			validPaths = append(validPaths, bfs...)
		} else if dfscounter < bfscounter {
			validPaths = append(validPaths, dfs...)
		} else {
			validPaths = append(validPaths, bfs...)
		}

	}
	return validPaths

}

//error checking for duplicate starting points
func PathDupeCheck(path [][]*Room) [][]*Room {

	dataMap := make(map[*Room][]*Room)

	for _, item := range path {
		if value, ok := dataMap[item[0]]; !ok {
			dataMap[item[0]] = item
		} else {
			if len(item) <= len(value) {
				dataMap[item[0]] = item

			}
		}
	}

	var output [][]*Room

	for _, value := range dataMap {
		output = append(output, value)
	}

	return output
}

//reassigns the slices in ascending (len) order
func Reassign(a [][]*Room) [][]*Room {

	sort.Slice(a, func(i, j int) bool {
		return len(a[i]) < len(a[j])
	})

	return a

}

func Min(a int, array [][]int) [][]int {
	result := array[0]
	for _, value := range array {

		if len(result) > len(value) {
			result = value
		}
	}
	result = append(result, a)
	return array

}

//returns a slice of slice with index 0 representing the number of rooms within a given path
func pathSlice(a [][]*Room) [][]int {
	var slice [][]int
	var s []int

	for i := range a {
		s = append(s, len(a[i]))
		slice = append(slice, s)
		s = []int{}
	}

	return slice
}

func pathMap(a [][]*Room) {
	pathmap := make(map[int][]int)

	for i := range a {
		pathmap[i] = append([]int{}, len(a[i]))

	}

	for key, value := range pathmap {
		fmt.Printf("%v----%v", key, value)
		fmt.Println()
	}

	fmt.Println(pathmap)

}

//finds most efficient path
func lowestInt(a [][]int, b [][]*Room) (int, []*Room) {

	min := 10000
	var path []*Room

	for i := 0; i < len(a); i++ {
		if a[i][0] < min {
			min = a[i][0]
			path = b[i]

		}

	}
	return min, path
}

func Increment(a [][]int, b int) [][]int {

	for _, slice := range a {
		if slice[0] == b {
			slice[0] += 1
			break
		}
	}
	return a

}

// Function to remove room using its key
func RemoveAnt(a []*Ant, b *Ant) []*Ant {
	ret := make([]*Ant, 0)
	if len(a) == 1 {
		return []*Ant{}
	}
	for i := 0; i < len(a); i++ {
		if a[i].key == b.key {
			ret = append(ret, a[:i]...)
			ret = append(ret, a[i+1:]...)
		}
	}
	return ret
}

func main() {

	bfsGraph := Graph{}

	//adding all rooms
	for i, line := range readAntsFile(args) {
		if strings.Contains(string(line), " ") {
			bfsGraph.AddRoom(strings.Split(readAntsFile(args)[i], " ")[0])
		}
		// adding all edges from and to rooms

		if strings.Contains(string(line), "-") {
			bfsGraph.AddEdge(strings.Split(readAntsFile(args)[i], "-")[0], strings.Split(readAntsFile(args)[i], "-")[1])
			bfsGraph.AddEdge(strings.Split(readAntsFile(args)[i], "-")[1], strings.Split(readAntsFile(args)[i], "-")[0])
		}

	}

	BFS(bfsGraph.getRoom(StartR), bfsGraph)
	//bfsGraph.Print()

	dfsGraph := Graph{}

	//adding all rooms
	for i, line := range readAntsFile(args) {
		if strings.Contains(string(line), " ") {
			dfsGraph.AddRoom(strings.Split(readAntsFile(args)[i], " ")[0])
		}
		// adding all edges from and to rooms
		// only adding edges in one direction to make the graph directional
		if strings.Contains(string(line), "-") {
			dfsGraph.AddEdge(strings.Split(readAntsFile(args)[i], "-")[0], strings.Split(readAntsFile(args)[i], "-")[1])
		}

	}

	DFS(dfsGraph.getRoom(StartR), dfsGraph)

	a := Ants{}

	Arrange := pathSlice(Reassign(PathDupeCheck(PathSelection(bfsPaths, dfsPaths))))
	Rooms := Reassign(PathDupeCheck(PathSelection(bfsPaths, dfsPaths)))

	// --------------------------------------------------------

	var unmovedAnts []*Ant
	var movedAnts []*Ant
	counter := 1

	for counter <= NumAnts(readAntsFile(args)) {

		number, _ := lowestInt(Arrange, Rooms)
		_, route := lowestInt(Arrange, Rooms)
		a.antz = append(a.antz, &Ant{key: "L" + strconv.Itoa(counter), path: route})
		Increment(Arrange, number)

		counter++
	}

	//-------------------------------------------------------------------
	// for _, paths := range validPaths {
	// 	for _, path := range paths {
	// 		fmt.Println(path.key)
	// 	}
	// }
	unmovedAnts = append(unmovedAnts, a.antz...)

	// for _ ant := range uunmovedAnts{

	// 	fmt.Println("Checking rooms in path: -----> ", ant)
	// }
	for len(unmovedAnts) > 0 || len(movedAnts) >= 1 {

		for _, ant := range unmovedAnts {
			if len(ant.path) == 1 {
				fmt.Print(ant.key, "-", ant.path[0].key, " ")
				ant.path[0].occupied = true
				unmovedAnts = RemoveAnt(unmovedAnts, ant)
				break
			}
		}

		// fmt.Println("FIRST  LOOP")
		for _, ant := range unmovedAnts {

			// fmt.Println()
			// if ant.path[0] == dfsGraph.getRoom(EndR) {
			// 	fmt.Print(ant.key, "-", ant.path[0].key, " ")
			// 	unmovedAnts = RemoveAnt(unmovedAnts, ant)
			// 	ant.path[0].occupied = false

			if !ant.path[0].occupied {
				fmt.Print(ant.key, "-", ant.path[0].key, " ")
				ant.path[0].occupied = true
				movedAnts = append(movedAnts, ant)
				unmovedAnts = RemoveAnt(unmovedAnts, ant)
				// fmt.Println(unmovedAnts)
			}
			// if len(ant.path) == 1 {
			// 	//fmt.Print(ant.key, "-+", ant.path[0].key, " ")
			// 	movedAnts = RemoveAnt(movedAnts, ant)
			// 	unmovedAnts = RemoveAnt(unmovedAnts, ant)
			// 	//ant.path[0].occupied = true
			// 	//continue
			// }
			// //bfsGraph.getRoom(EndR).occupied = true
			// //	dfsGraph.getRoom(EndR).occupied = true

			// ant.path[0].occupied = false
			//ant.path = ant.path[1:]

		}

		fmt.Println()

		// fmt.Println("SECOND LOOP")
		for _, ant := range movedAnts {

			// fmt.Println("moving", ant.key)
			// movedAnts = RemoveAnt(movedAnts, ant)

			if len(ant.path) > 1 {
				ant.path[0].occupied = false

				ant.path = ant.path[1:]
				fmt.Print(ant.key, "-", ant.path[0].key, " ")
				// not sure what is happening with this if statement
				// if len(ant.path) > 1 {
				// 	continue
			} else {
				movedAnts = RemoveAnt(movedAnts, ant)
				ant.path = []*Room{}
			}
		}

	}

	fmt.Println()

	// fmt.Print("Unmoved Ants", " : ")
	// for _, value := range unmovedAnts {
	// 	fmt.Print(value.key)
	// 	fmt.Print("  ")
	// }
	// fmt.Println()

	// fmt.Print("Moved Ants", " : ")
	// for _, value := range movedAnts {
	// 	fmt.Print(value.key)
	// 	fmt.Print("  ")
	// }

	// fmt.Println()

}
