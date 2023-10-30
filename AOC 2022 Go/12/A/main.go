package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	x, y int
}

type Queue struct {
	elements []Coord
}

func (q *Queue) enqueue(coord Coord) {
	q.elements = append(q.elements, coord)
}

func (q *Queue) dequeue() Coord {
	out := q.elements[0]
	q.elements = q.elements[1:]
	return out
}

var (
	grid [][]byte
	min  map[Coord]int = make(map[Coord]int, 0)
	que  Queue
)

// use breath first search to find most optimal path
func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		grid = append(grid, []byte(line))
	}

out:
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				min[Coord{i, j}] = 0
				que.enqueue(Coord{i, j})
				grid[i][j] = 'a'
				break out
			}
		}
	}

	for len(que.elements) != 0 {
		coord := que.dequeue()
		x := coord.x
		y := coord.y

		// Up
		if x > 0 {
			if grid[x-1][y]-1 <= grid[x][y] {
				_, ok := min[Coord{x - 1, y}]
				if !ok {
					min[Coord{x - 1, y}] = min[coord] + 1
					que.enqueue(Coord{x - 1, y})
				}
			}

			if grid[x-1][y] == 'E' && grid[x][y] >= 'y' {
				fmt.Println(min[coord] + 1)
				break
			}
		}

		// Down
		if x < len(grid)-1 {
			if grid[x+1][y]-1 <= grid[x][y] {
				_, ok := min[Coord{x + 1, y}]
				if !ok {
					min[Coord{x + 1, y}] = min[coord] + 1
					que.enqueue(Coord{x + 1, y})
				}
			}

			if grid[x+1][y] == 'E' && grid[x][y] >= 'y' {
				fmt.Println(min[coord] + 1)
				break
			}
		}

		// Left
		if y > 0 {
			if grid[x][y-1]-1 <= grid[x][y] {
				_, ok := min[Coord{x, y - 1}]
				if !ok {
					min[Coord{x, y - 1}] = min[coord] + 1
					que.enqueue(Coord{x, y - 1})
				}
			}

			if grid[x][y-1] == 'E' && grid[x][y] >= 'y' {
				fmt.Println(min[coord] + 1)
				break
			}
		}

		// Up
		if y < len(grid[0])-1 {
			if grid[x][y+1]-1 <= grid[x][y] {
				_, ok := min[Coord{x, y + 1}]
				if !ok {
					min[Coord{x, y + 1}] = min[coord] + 1
					que.enqueue(Coord{x, y + 1})
				}
			}

			if grid[x][y+1] == 'E' && grid[x][y] >= 'y' {
				fmt.Println(min[coord] + 1)
				break
			}
		}

	}

	// for i := range grid {
	// 	for j := range grid[i] {
	// 		v, k := min[Coord{i, j}]
	// 		if k {
	// 			fmt.Printf("[%3d]", v)
	// 		} else {
	// 			fmt.Print("[   ]")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
}
