package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

type Valve struct {
	ID   string
	Flow int

	Connection []string
}

var (
	valveIdx   map[string]int = make(map[string]int)
	valves     []Valve        = make([]Valve, 0)
	connection [][]int        = make([][]int, 0)

	max int = 0
)

func printConnection() {
	fmt.Printf("  ")
	for k := range valves {
		fmt.Printf(" %v ", valves[k].ID)
	}
	fmt.Println()
	for from := range valves {
		fmt.Printf("%v", valves[from].ID)
		for to := range valves {
			fmt.Printf("[%2d]", connection[from][to])
		}
		fmt.Println()
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	idx := 0
	for sc.Scan() {
		var valve Valve
		lines := strings.Split(sc.Text(), "; ")
		fmt.Sscanf(lines[0], "Valve %s has flow rate=%d", &valve.ID, &valve.Flow)

		conns := strings.Split(lines[1][22:], ", ")
		valve.Connection = make([]string, 0)
		for i := range conns {
			valve.Connection = append(valve.Connection, strings.TrimSpace(conns[i]))
		}

		valves = append(valves, valve)
		valveIdx[valve.ID] = idx
		idx++
	}

	for from := range valves {
		connection = append(connection, make([]int, len(valves)))
		routes := make([]string, 0)
		length := 1

		routes = append(routes, valves[from].Connection...)

		for len(routes) != 0 {
			newRoute := make([]string, 0)
			// Check all possible route
			for j := range routes {
				to := valveIdx[routes[j]]
				if connection[from][to] == 0 {
					connection[from][to] = length
					newRoute = append(newRoute, valves[to].Connection...)
				}
			}
			length++
			routes = newRoute
		}
	}

	for i := 0; i < len(valves); {
		if valves[i].ID != "AA" && valves[i].Flow == 0 {
			valves = append(valves[:i], valves[i+1:]...)
			connection = append(connection[:i], connection[i+1:]...)
			for j := range connection {
				connection[j] = append(connection[j][:i], connection[j][i+1:]...)
			}
		} else {
			i++
		}
	}

	valveIdx = make(map[string]int)
	for i := range valves {
		valveIdx[valves[i].ID] = i
	}

	printConnection()

	path := make([]int, 0)
	for i := range valves {
		if valves[i].ID == "AA" {
			path = append(path, i)
			break
		}
	}

	getAnswer(path, 30, 0)
	fmt.Println(max)
}

func getAnswer(path []int, minutes, currentMax int) {
	if minutes <= 0 {
		if currentMax > max {
			max = currentMax
			fmt.Println(path, max)
		}
		return
	}
	currentValve := &valves[path[len(path)-1]]
	currentMax += currentValve.Flow * minutes
	minutes--

	nextPaths := make([]int, 0)
	for i := range valves {
		if !slices.Contains(path, i) {
			nextPaths = append(nextPaths, i)
		}
	}

	if len(nextPaths) == 0 {
		if currentMax > max {
			max = currentMax
			fmt.Println(path, max)
		}
	}

	for i := range nextPaths {
		path = append(path, nextPaths[i])
		from := path[len(path)-2]
		to := path[len(path)-1]
		travelTime := connection[from][to]

		getAnswer(path, minutes-travelTime, currentMax)

		path = path[:len(path)-1]
	}
}

// DD BB JJ HH EE CC
// 0 3 1 6 5 4 2
