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
	s = readAntsFile("ants.txt")
	if s[0] <= "0" {
		err := fmt.Errorf("invalid number of ants")
		fmt.Println(err.Error())
	}

	antNumInt, _ := strconv.Atoi(antNum)
	return antNumInt
}

// Gets out the start room and returns it
func StartRoom([]string) string {

	var startRoom string
	s := readAntsFile("ants.txt")

	for i := 0; i < len(s); i++ {
		if s[i] == "##start" {
			startRoom = strings.Split(string(s[i+1]), " ")[0]
		}

	}

	return startRoom

}

// Gets out the end room and returns it
func EndRoom([]string) string {
	var endRoom string
	s := readAntsFile("ants.txt")

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

func main() {

	bfsGraph := Graph{}

	//adding all rooms
	for i, line := range readAntsFile("ants.txt") {
		if strings.Contains(string(line), " ") {
			bfsGraph.AddRoom(strings.Split(readAntsFile("ants.txt")[i], " ")[0])
		}
		// adding all edges from and to rooms

		if strings.Contains(string(line), "-") {
			bfsGraph.AddEdge(strings.Split(readAntsFile("ants.txt")[i], "-")[0], strings.Split(readAntsFile("ants.txt")[i], "-")[1])
			bfsGraph.AddEdge(strings.Split(readAntsFile("ants.txt")[i], "-")[1], strings.Split(readAntsFile("ants.txt")[i], "-")[0])
		}

	}

	//bfsGraph.Print()
	BFS(bfsGraph.getRoom(StartR), bfsGraph)
	//bfsGraph.Print()

	dfsGraph := Graph{}

	//adding all rooms
	for i, line := range readAntsFile("ants.txt") {
		if strings.Contains(string(line), " ") {
			dfsGraph.AddRoom(strings.Split(readAntsFile("ants.txt")[i], " ")[0])
		}
		// adding all edges from and to rooms
		// only adding edges in one direction to make the graph directional
		if strings.Contains(string(line), "-") {
			dfsGraph.AddEdge(strings.Split(readAntsFile("ants.txt")[i], "-")[0], strings.Split(readAntsFile("ants.txt")[i], "-")[1])
		}

	}

	//dfsGraph.Print()
	DFS(dfsGraph.getRoom(StartR), dfsGraph)

	for _, r := range validPaths {
		for _, f := range r {
			fmt.Println(f.key)
		}
	}

	//dfsGraph.Print()

	//Output()

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
// Add edge to the graph. deals with a directional graph only but condition in the main makes it undirected
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

				//fmt.Println("*", vList)
				nbr.path = append(r.path, nbr)
				if contains(nbr.path, EndR) {
					for _, r := range nbr.path {
						fmt.Println("DFS Finale: ----->", r.key)
					}
					validPaths = append(validPaths, nbr.path)

				}
				//fmt.Println(nbr.path)
				vList = append(vList, nbr.key)
				//fmt.Println("La Final1: ", vList)

				//DeleteEdge(r, g)
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

				//fmt.Println("*", vList)
				nbr.path = append(r.path, nbr)
				if contains(nbr.path, EndR) {
					fmt.Println("dfs print check", nbr.path)

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

	// for _, v := range queue {
	// 	fmt.Println("P2", v.key, "\t")
	// 	fmt.Println("check --------========", queue[0].adjacent[0].key)

	// }

	// checks if there is a link between start and end directly
	for i, v := range g.getRoom(StartR).adjacent {
		if v.key == g.getRoom(EndR).key {
			g.getRoom(EndR).path = append(g.getRoom(EndR).path, g.getRoom(StartR))
			vPaths = append(vPaths, g.getRoom(EndR).path)
			g.getRoom(StartR).adjacent = append(g.getRoom(StartR).adjacent[:i], g.getRoom(StartR).adjacent[i+1:]...)
			//fmt.Println("End reached --------------------------------------------------------------------->", g.getRoom(StartR).path)
		}

	}

	//fmt.Println("QQQ:", queue[0].key)

	//checks the queue for the end room and if the queue is not empty

	for !contains(queue, g.getRoom(EndR).key) && len(queue) >= 1 {
		qfront := queue[0]
		//	fmt.Println("QF:", qfront.key)

		for _, room := range qfront.adjacent {
			if !room.visited {
				room.visited = true
				room.path = append(qfront.path, room)
				//	fmt.Println("QFront:\n", qfront.key)
				//	fmt.Println("Path:", room.path)
				queue = append(queue, room)
				// fmt.Println(queue)
			}
			// for _, v := range queue {
			// 	fmt.Print("print check +++++", v.key, "\t")
			// }

		}
		// fmt.Println("YYYYYYY")
		// fmt.Println(r.key)
		// fmt.Println(g.getRoom(StartR).key)
		// fmt.Println("YYYYYYY")
		queue = queue[1:]
		// fmt.Println("OOOOOOO")
		// fmt.Println(r.key)
		// fmt.Println(g.getRoom(StartR).key)
		// fmt.Println("OOOOOOO")

		//checking if the end room has been queued/reached
		if doesContainRoom(queue, g.getRoom(EndR).key) {

			for _, room := range g.rooms {
				room.visited = false
			}
			//	fmt.Println("End reached--------------------------------------------------------------------->:", qfront.path)
			vPaths = append(vPaths, qfront.path)

			//fmt.Println("team end check: ", vPaths)

			DeleteEdge(qfront, g)
			//	fmt.Println("#### LEVEL TEST ####")
			//	fmt.Println(len(r.adjacent))
			// for _, v := range r.adjacent {
			// 	fmt.Println(v.key, "\t")
			// }
			//	fmt.Println("len of adj list", len(g.getRoom(StartR).adjacent))

			if len(g.getRoom(StartR).adjacent) == 0 {
				//	fmt.Println("loop 5")
				break
			}
			//g.Print()

			if len(g.getRoom(StartR).adjacent) >= 1 {
				for _, froom := range g.getRoom(StartR).adjacent {
					for _, sroom := range froom.adjacent {
						if sroom.key != g.getRoom(EndR).key {
							//	fmt.Println("loop 8")
							break
						} else {
							//	fmt.Println("loop 9")
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
		fmt.Printf("BFS Finale: ----> ")
		for _, r := range v {
			fmt.Printf("%v ", r.key)
		}
		fmt.Println()
	}

}

// func Output() {
// 	numOfAnts := NumAnts(readAntsFile("ants.txt"))
// 	// valid paths from dfs function
// 	unmovedAnts := []string{}

// 	for i := 1; i <= numOfAnts; i++ {
// 		unmovedAnts = append(unmovedAnts, strconv.Itoa(i))
// 	}
// 	// fmt.Println(unmovedAnts)
// 	vp := validPaths

// 	// map to hold each room with visited bool
// 	// append each path into the map
// 	occupied := make(map[string]bool)

// 	// for _, pathslice := range vp {
// 	// 	for _, room := range pathslice {
// 	// 		occupied[room] = false
// 	// 	}
// 	// }

// 	movingAnts := []string{}
// 	counter := 0
// 	// if room is unoccupied, add ant into moving slice
// 	// if all rooms upto that point are occupied, turn done, println
// 	// for _, path := range vp {
// 	// 	for _, room := range path {
// 	// 		if !occupied[room] {
// 	// 			movingAnts = append(movingAnts, unmovedAnts[0])
// 	// 			if len(unmovedAnts) > 1 {
// 	// 				unmovedAnts = unmovedAnts[1:]
// 	// 			}

// 				//fmt.Println("checking mAnts: --->", movingAnts)
// 				for i, ant := range movingAnts {
// 					if room == EndR {
// 						movingAnts = movingAnts[1:]
// 						fmt.Print("check", movingAnts)
// 					}
// 					fmt.Printf("L%v-%v ", ant, room)
// 					occupied[room] = true
// 					if room != path[0] && i == len(movingAnts)-1 {
// 						//fmt.Println(v)
// 						room = path[0]
// 					}

// 				}
// 				occupied[room] = false
// 				// fmt.Println(occupied)
// 				// for turn
// 				fmt.Println()
// 			}
// 			// } else if !occupied[room] {
// 			// 	movingAnts = append(movingAnts, unmovedAnts[0])
// 			// 	unmovedAnts = unmovedAnts[1:]
// 			// 	// occupied[path[v+1]] = true
// 			// 	for _, ant := range movingAnts {
// 			// 		fmt.Printf("L%v-%v ", ant, room)
// 			// 		occupied[room] = true
// 			// 		// fmt.Println(occupied)
// 			// 		occupied[path[v-1]] = false
// 			// 		room = path[v-1]
// 			// 		// fmt.Println(occupied)
// 			// 		// if ant goes to the end room, remove it from moving ants
// 			// 		if occupied[room] && room == EndR {
// 			// 			movingAnts = movingAnts[1:]
// 			// 			fmt.Print(movingAnts)
// 			// 		}
// 			// 	}
// 			// 	occupied[room] = false
// 			// 	fmt.Println()
// 			// }
// 			//fmt.Println("2nd checking mAnts: +++++>", movingAnts)
// 			//counter++
// 		//}
// 		// for r, v := range occupied {
// 		// 	fmt.Println("Printing occupied map: --> ", r, v)
// 		// }
// 	//}
// }
