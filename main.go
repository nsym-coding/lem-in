package main

import (
	"bufio"
	//"container/list"
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
	path     []string
	visited  bool
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

//delete edge from starting room
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

	//adding all rooms
	for i, line := range readAntsFile("ants.txt") {
		if strings.Contains(string(line), " ") {
			test.AddRoom(strings.Split(readAntsFile("ants.txt")[i], " ")[0])
		}
		// adding all edges from and to rooms
		// maybe add a condition so that it adds the edges in order i.e. the end room as the last edge?
		if strings.Contains(string(line), "-") {
			test.AddEdge(strings.Split(readAntsFile("ants.txt")[i], "-")[0], strings.Split(readAntsFile("ants.txt")[i], "-")[1])
			test.AddEdge(strings.Split(readAntsFile("ants.txt")[i], "-")[1], strings.Split(readAntsFile("ants.txt")[i], "-")[0])
		}

	}

	test.Print()
	//DFS(test.getRoom(StartR), test)
	//dfsStart(test)
	//GFS(test)
	//test.PrintPath()
	//DeleteEdge()
	BFS(test.getRoom(StartR), test)
	//BreadthFS(test.getRoom(StartR), &test)
	test.Print()

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
		// toRoom.adjacent = append(toRoom.adjacent, fromRoom)
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

// Depth first search function that operates recursively
func DFS(r *Room, g Graph) {

	//vList := []string{}
	sRoom := g.getRoom(StartR)

	// set the room being checked visited status to true
	if r.key != EndR {
		r.visited = true

		// append the r key to the visited list
		//vList = append(vList, r.key)
		// range through the neighbours of the r
		for _, nbr := range r.adjacent {
			if !nbr.visited {
				/* for each neighbour that hasn't been visited,
				- append their key to the visited slice,
				- then apply dfs to them recursively,
				- then append their key to their path value
				*/

				//fmt.Println("*", vList)
				nbr.path = append(r.path, nbr.key)
				if doesContain(EndR, nbr.path) {
					fmt.Println("dfs printing path", nbr.path)
				}
				//fmt.Println(nbr.path)
				//vList = append(vList, nbr.key)
				DFS(nbr, Graph{g.rooms})

			}

		}

	} else {
		if len(sRoom.adjacent) > 1 && !contains(sRoom.adjacent, EndR) {
			// vList = append(vList, r.key)
			//fmt.Println("*", vList)
			sRoom.adjacent = sRoom.adjacent[1:]

			DFS(sRoom, Graph{g.rooms})

			// } else {
			// 	// vList = append(vList, r.key)
			// 	//fmt.Println("*", vList)
			// }
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

		// append the r key to the visited list
		//vList = append(vList, r.key)
		// range through the neighbours of the r
		for _, nbr := range r.adjacent {
			if !nbr.visited {
				/* for each neighbour that hasn't been visited,
				- append their key to the visited slice,
				- then apply dfs to them recursively,
				- then append their key to their path value
				*/

				//fmt.Println("*", vList)
				nbr.path = append(r.path, nbr.key)
				if doesContain(EndR, nbr.path) {
					fmt.Println("dfs print check",nbr.path)
					//fmt.Println("nbr length in func -------", len(nbr.adjacent))
				
					
					//g.Print()
					return true
					//BFS(nbr, g)
					//return true
				}
				//fmt.Println(nbr.path)
				//vList = append(vList, nbr.key)
				//DFSBFS(nbr, Graph{g.rooms})

			}

		}

	// } else {
	// 	// if len(sRoom.adjacent) > 1 && !contains(sRoom.adjacent, EndR) {
	// 		// vList = append(vList, r.key)
	// 		//fmt.Println("*", vList)

	// 		return true
	// 		//sRoom.adjacent = sRoom.adjacent[1:]

	// 	//	DFS(sRoom, Graph{g.rooms})

	// 		// } else {
	// 		// 	// vList = append(vList, r.key)
	// 		// 	//fmt.Println("*", vList)
	// 		// }
	// 	}
}
return false
}


// Function that initialises that DFS algorithm by taking the target graph as an argument
// func dfsStart(g *Graph) {
// 	DFS(g.getRoom(StartR))
// }
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
				if room.adjacent[j].key == r.path[i] {
					room.adjacent = remove(room.adjacent, r.path[i])
				}
			}
		}
	}
}

func BFS(r *Room, g Graph) {

	var vPaths [][]string

	//sRoom := g.getRoom(StartR)
	//queue variable, procedurally populated with rooms yet to be visited
	var queue []*Room

	//set start room as visited
	r.visited = true

	//initialise queue with start room
	queue = append(queue, r)

	for _, v := range queue {
		fmt.Println("P2", v.key, "\t")
		fmt.Println("check --------========", queue[0].adjacent[0].key)

	}

	// checks if there is a link between start and end directly
	for i, v := range g.getRoom(StartR).adjacent {
		if v.key == g.getRoom(EndR).key {
			g.getRoom(EndR).path = append(g.getRoom(EndR).path, g.getRoom(StartR).key)
			vPaths = append(vPaths, g.getRoom(EndR).path)
			g.getRoom(StartR).adjacent = append(g.getRoom(StartR).adjacent[:i], g.getRoom(StartR).adjacent[i+1:]...)
			fmt.Println("End reached --------------------------------------------------------------------->", g.getRoom(StartR).path)
		}
		//continue
	}

	fmt.Println("QQQ:", queue[0].key)
	//fmt.Println("Queue", queue)
	//checks the queue for a non-zero value
	// for len(queue) > 0 {
	for !contains(queue, g.getRoom(EndR).key) && len(queue) >= 1{
		qfront := queue[0]
		fmt.Println("QF:", qfront.key)

		for _, room := range qfront.adjacent {
			if !room.visited {
				room.visited = true
				room.path = append(qfront.path, room.key)
				fmt.Println("QFront:\n", qfront.key)
				fmt.Println("Path:", room.path)
				queue = append(queue, room)
				// fmt.Println(queue)
			}
			for _, v := range queue {
				fmt.Print("print check +++++", v.key, "\t")
			}

			// if checkEnd(g, room) {
			// 	fmt.Println(room.path)
			// 	os.Exit(0)
			// }
		}
		fmt.Println("YYYYYYY")
		fmt.Println(r.key)
		fmt.Println(g.getRoom(StartR).key)
		fmt.Println("YYYYYYY")
		queue = queue[1:]
		fmt.Println("OOOOOOO")
		fmt.Println(r.key)
		fmt.Println(g.getRoom(StartR).key)
		fmt.Println("OOOOOOO")

		//checking if the end room has been queued/reached
		if doesContainRoom(queue, g.getRoom(EndR).key) {
			
			//DeleteEdge(qfront, g)
			//fmt.Println("Queue when end reached")
			for _, room := range g.rooms {
				room.visited = false
			}
			fmt.Println("End reached--------------------------------------------------------------------->:", qfront.path)
			vPaths = append(vPaths, qfront.path)
			
			fmt.Println("team end check: ", vPaths)

			//iterating through start room's adjacents and removing the lead room
			// for i := 0; i < len(r.adjacent); i++ {
			// 	if r.adjacent[i].key == qfront.path[0] {
			// 		//fmt.Println("First home:", r.adjacent[i].key, "\t")
			// 		r.adjacent = append(r.adjacent[:i], r.adjacent[i+1:]...)

			// 		// fmt.Println("Start's adjacents:", r.adjacent)
			// 		for _, v := range r.adjacent {
			// 			fmt.Print("print check",v.key, "\t")
			// 		}
			// 		fmt.Println()
			// 	}
			// 	fmt.Println("NNNNNN")
			// 	fmt.Println(r.key)
			// 	fmt.Println(g.getRoom(StartR).key)
			// 	fmt.Println("NNNNNN")

			// 	// for _, lnks := range qfront.adjacent{
			// 		// 	DeleteEdge(lnks, g)
			// 	// }

			// 	// fmt.Println("test1", r.adjacent[i].key)
			// 	// if queue[0] == g.getRoom(StartR) && len(queue) == 1 {
			// 		// 	BFS(g.getRoom(StartR), Graph{g.rooms}, queue)
			// 		// } else {
			// 			//queue = queue[1:]
			// 			// 	BFS(g.getRoom(StartR), Graph{g.rooms}, queue)
			// 			// }
			// 		}
					DeleteEdge(qfront, g)
			fmt.Println("#### LEVEL TEST ####")
			fmt.Println(len(r.adjacent))
			for _, v := range r.adjacent {
				fmt.Println(v.key, "\t")
			}
			fmt.Println("len of adj list", len(g.getRoom(StartR).adjacent))
			//if len(g.getRoom(StartR).adjacent) == 1 {
			// for _, r := range g.getRoom(StartR).adjacent {
			// 	if len(r.adjacent) == 0 {
			// 		break
			// 	}
			// }
			if len(g.getRoom(StartR).adjacent) == 0 {
				fmt.Println("loop 5")
				break
			} 
			g.Print()
			
			if len(g.getRoom(StartR).adjacent) >= 1 {
				for _, froom := range g.getRoom(StartR).adjacent {
					for _, sroom := range froom.adjacent {
						if sroom.key != g.getRoom(EndR).key {
							fmt.Println("loop 8")
							break
						} else {
							fmt.Println("loop 9")
							BFS(g.getRoom(StartR), Graph{g.rooms})
							queue = queue[1:]
						}
					}
				}
			}
			BFS(g.getRoom(StartR), Graph{g.rooms})
	// 		if len(g.getRoom(StartR).adjacent) >= 1 {
	// 			fmt.Println("start has 1 ------")
	// 			fmt.Println("len check in loop ----------------------",len(g.getRoom(StartR).adjacent))
	// 			counter := 0 
				
				

	// 			for _, value := range g.getRoom(StartR).adjacent{
	// 				fmt.Println("HELLOOOOOOOOOOOOOOOOOOOO!")
	// 				if DFSBFS(value, Graph{g.rooms}) == true{
						
	// 					counter ++
	// 					fmt.Println("checking culprit---------------+++++++++++++++++++++++++++++++++++")
	// 				}
	// 				fmt.Println("counter check YYYYYYYYYYYYY", counter)
	// 			}

	// 			fmt.Println("counter check +++++++++++++", counter)
	// 			if counter != 0{
		
	// 		BFS(g.getRoom(StartR), Graph{g.rooms})
	// 		//queue = queue[1:]
	// 	}
	// }
}
}
for _, v := range vPaths {
	fmt.Println("Finale: ", v)
}
//mt.Println("path check ----------------------", vPaths)
}

// fmt.Println("\nQUEUE BEFORE")
// for _, v := range queue {
// 	fmt.Print(v.key)
// 	}
// 	fmt.Println()
//queue = queue[1:]
// 	fmt.Println("\nQUEUE AFTER")
// 	for _, v := range queue {
// 	fmt.Print(v.key)
// 	}
// 	fmt.Println()

func BreadthFS(r *Room, g *Graph) {

	var vPath [][]string

	queue := []*Room{r}

	for len(queue) > 0 {
		//queue = queue[1:]

		// checks if there is a link between start and end directly
		for i, v := range g.getRoom(StartR).adjacent {
			if v.key == g.getRoom(EndR).key {
				g.getRoom(EndR).path = append(g.getRoom(EndR).path, g.getRoom(StartR).key)
				vPath = append(vPath, g.getRoom(EndR).path)
				g.getRoom(StartR).adjacent = append(g.getRoom(StartR).adjacent[:i], g.getRoom(StartR).adjacent[i+1:]...)
				fmt.Println("loop 1")
			}
			//continue
		}

		current := queue[0]
		queue = queue[1:]
		for !contains(queue, g.getRoom(EndR).key) {
			for _, nbr := range r.adjacent {
				if !nbr.visited {
					nbr.visited = true
					nbr.path = append(current.path, nbr.key)
					queue = append(queue, nbr)
					fmt.Println("loop 2")
					//	continue

				}
			}
		}
		if contains(queue, g.getRoom(EndR).key) {
			for _, room := range g.rooms {
				room.visited = false
				fmt.Println("loop 3")
			}
			vPath = append(vPath, current.path)
			for i := 0; i < len(g.getRoom(StartR).adjacent); i++ {
				if r.adjacent[i].key == current.path[0] {
					r.adjacent = append(r.adjacent[:i], r.adjacent[i+1:]...)
					fmt.Println("loop 4")
				}
			}
			DeleteEdge(current, Graph{}) //maybe put this after loop below

			if len(g.getRoom(StartR).adjacent) == 0 {
				fmt.Println("loop 5")
				break
			} else if len(g.getRoom(StartR).adjacent) == 1 {
				for _, v := range g.getRoom(StartR).adjacent[0].adjacent {
					if v.key != g.getRoom(EndR).key {
						fmt.Println("loop 6")
						break
					} else {
						fmt.Println("loop 7")
						BreadthFS(g.getRoom(StartR), g)
					}
				}
			} else if len(g.getRoom(StartR).adjacent) >= 2 {
				for _, froom := range g.getRoom(StartR).adjacent {
					for _, sroom := range froom.adjacent {
						if sroom.key != g.getRoom(EndR).key {
							fmt.Println("loop 8")
							break
						} else {
							fmt.Println("loop 9")
							BreadthFS(g.getRoom(StartR), g)
							queue = queue[1:]
						}
					}
				}
			}

		}
	}
	for _, r := range vPath {
		fmt.Println("Finale: ", r)
	}

}

// Delete all the edges and the rooms once a path to the end has been found that uses them
// Changed the condition so that when we build double links, we make sure there's no double links with start
// Need to make sure that doesn't affect the other graphs where the link goes from the end and towards start
