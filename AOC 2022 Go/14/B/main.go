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
	// it takes too long time too completely fill out the whole triangle, take the height of what would be right and left triangle to calculate them instead of counting them
	minLeft, minRight int
)

func setMinMax(c Coord) {
	if c.x-1 < minX {
		minX = c.x - 1
	} else if c.x+1 > maxX {
		maxX = c.x + 1
	}

	if c.y+1 > maxY {
		maxY = c.y + 1
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
	minLeft, minRight = maxY, maxY

	for {
		sand := Coord{500, 0}
		if _, k := grid[sand]; k {
			break
		}

		for {
			// floor
			if sand.y+1 > maxY {
				grid[sand] = 'O'
				cnt++
				break
			}
			// down
			if _, k := grid[Coord{sand.x, sand.y + 1}]; !k {
				sand.y++
				continue
			}

			// left
			if sand.x-1 < minX {
				if minLeft > sand.y+1 {
					minLeft = sand.y + 1
				}
			} else {
				if _, k := grid[Coord{sand.x - 1, sand.y + 1}]; !k {
					sand.x--
					sand.y++
					continue
				}
			}

			// right wall
			if sand.x+1 > maxX {
				if minRight > sand.y+1 {
					minRight = sand.y + 1
				}
			} else {
				if _, k := grid[Coord{sand.x + 1, sand.y + 1}]; !k {
					sand.x++
					sand.y++
					continue
				}
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
				continue
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

	fmt.Println(maxY-minLeft, maxY-minRight)
	leftTriangle := (maxY - minLeft + 2) * (maxY - minLeft + 1) / 2
	rightTriangle := (maxY - minRight + 2) * (maxY - minRight + 1) / 2
	fmt.Println(cnt, leftTriangle, rightTriangle)
	fmt.Println(cnt + leftTriangle + rightTriangle)

}
