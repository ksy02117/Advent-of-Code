package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

var (
	grid map[Coord]byte = make(map[Coord]byte)

	minX, maxX int = 500, 500
	maxY       int = 0
)

// sets boundary since if sand ever reaches out of bound on left, right, or bottom, it will fall indefinitely
func setMinMax(c Coord) {
	if c.x < minX {
		minX = c.x
	} else if c.x > maxX {
		maxX = c.x
	}

	if c.y > maxY {
		maxY = c.y
	}
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		coordinates := strings.Split(sc.Text(), " -> ")
		var prev, next Coord
		fmt.Sscanf(coordinates[0], "%d,%d", &prev.x, &prev.y)
		setMinMax(prev)

		for i := 1; i < len(coordinates); i++ {
			fmt.Sscanf(coordinates[i], "%d,%d", &next.x, &next.y)
			setMinMax(next)

			if prev.x == next.x {
				var high, low int
				if prev.y < next.y {
					high, low = next.y, prev.y
				} else {
					high, low = prev.y, next.y
				}

				for i := low; i <= high; i++ {
					grid[Coord{prev.x, i}] = '#'
				}
			}

			if prev.y == next.y {
				var high, low int
				if prev.x < next.x {
					high, low = next.x, prev.x
				} else {
					high, low = prev.x, next.x
				}

				for i := low; i <= high; i++ {
					grid[Coord{i, prev.y}] = '#'
				}
			}
			prev = next
		}
	}

	cnt := 0
outer:
	for {
		sand := Coord{500, 0}

		for {
			// down
			if sand.y+1 > maxY {
				break outer
			}
			if _, k := grid[Coord{sand.x, sand.y + 1}]; !k {
				sand.y++
				continue
			}

			// left
			if sand.x-1 < minX {
				break outer
			}
			if _, k := grid[Coord{sand.x - 1, sand.y + 1}]; !k {
				sand.x--
				sand.y++
				continue
			}

			// right
			if sand.x+1 > maxX {
				break outer
			}
			if _, k := grid[Coord{sand.x + 1, sand.y + 1}]; !k {
				sand.x++
				sand.y++
				continue
			}

			grid[sand] = 'O'
			cnt++
			break
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if x == 500 && y == 0 {
				fmt.Print("X")
			}
			v, k := grid[Coord{x, y}]
			if k {
				fmt.Printf("%c", v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println(minX, maxX, maxY)
	fmt.Println(cnt)

}
