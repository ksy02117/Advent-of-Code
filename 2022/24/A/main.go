package main

import (
	"bufio"
	"fmt"
	"os"
)

type Player struct {
	x, y int
}

type Blizzard struct {
	x, y int
	dir  int
}

func (b *Blizzard) Move() {
	switch b.dir {
	case 0: // Right
		if b.x == len(grid[0])-2 {
			b.x = 1
		} else {
			b.x++
		}
	case 1: // Down
		if b.y == len(grid)-2 {
			b.y = 1
		} else {
			b.y++
		}
	case 2: // Left
		if b.x == 1 {
			b.x = len(grid[0]) - 2
		} else {
			b.x--
		}
	case 3: // Up
		if b.y == 1 {
			b.y = len(grid) - 2
		} else {
			b.y--
		}
	}
}

func isSafe(x, y int) bool {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) {
		return false
	}
	if grid[y][x] == '#' {
		return false
	}
	for i := range blizzards {
		if blizzards[i].x == x && blizzards[i].y == y {
			return false
		}
	}

	return true
}

func printMap() {
	for y := range grid {
		for x := range grid[y] {
			cnt := 0
			var tmp Blizzard
			for i := range blizzards {
				if blizzards[i].x == x && blizzards[i].y == y {
					cnt++
					tmp = blizzards[i]
				}
			}

			if cnt == 0 {
				fmt.Print(string(grid[y][x]))
			} else if cnt != 1 {
				fmt.Print(cnt)
			} else {
				switch tmp.dir {
				case 0:
					fmt.Print(">")
				case 1:
					fmt.Print("v")
				case 2:
					fmt.Print("<")
				case 3:
					fmt.Print("^")
				}
			}
		}
		fmt.Println()
	}
}

var (
	grid      [][]byte   = make([][]byte, 0)
	blizzards []Blizzard = make([]Blizzard, 0)

	playersBuffer map[Player]bool = make(map[Player]bool)
	players       map[Player]bool
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for y := 0; sc.Scan(); y++ {
		line := sc.Text()

		row := make([]byte, 0)
		for x := 0; x < len(line); x++ {
			switch line[x] {
			case '>':
				blizzards = append(blizzards, Blizzard{x, y, 0})
				row = append(row, '.')
			case 'v':
				blizzards = append(blizzards, Blizzard{x, y, 1})
				row = append(row, '.')
			case '<':
				blizzards = append(blizzards, Blizzard{x, y, 2})
				row = append(row, '.')
			case '^':
				blizzards = append(blizzards, Blizzard{x, y, 3})
				row = append(row, '.')
			default:
				row = append(row, line[x])
			}
		}

		grid = append(grid, row)
	}

	// Player Location
	playersBuffer[Player{1, 0}] = true

	// fmt.Println("Minute :", 0)
	// printMap()

	minute := 1
Outer:
	for {
		// Move Blizzard
		for i := range blizzards {
			blizzards[i].Move()
		}

		// fmt.Println("Minute :", minute)
		// printMap()

		// Move Player
		players = playersBuffer
		playersBuffer = make(map[Player]bool)
		for coord := range players {
			// Right
			if isSafe(coord.x+1, coord.y) {
				playersBuffer[Player{coord.x + 1, coord.y}] = true
			}

			// Down
			if isSafe(coord.x, coord.y+1) {
				if coord.y+1 == len(grid)-1 {
					minute++
					break Outer
				}
				playersBuffer[Player{coord.x, coord.y + 1}] = true
			}

			// Left
			if isSafe(coord.x-1, coord.y) {
				playersBuffer[Player{coord.x - 1, coord.y}] = true
			}

			// Up
			if isSafe(coord.x, coord.y-1) {
				playersBuffer[Player{coord.x, coord.y - 1}] = true
			}

			// Stay
			if isSafe(coord.x, coord.y) {
				playersBuffer[Player{coord.x, coord.y}] = true
			}
		}

		minute++
	}

	playersBuffer = make(map[Player]bool)
	playersBuffer[Player{len(grid[0]) - 2, len(grid) - 1}] = true
Outer2:
	for {
		// Move Blizzard
		for i := range blizzards {
			blizzards[i].Move()
		}

		// fmt.Println("Minute :", minute+1)
		// printMap()

		// Move Player
		players = playersBuffer
		playersBuffer = make(map[Player]bool)
		for coord := range players {
			// Right
			if isSafe(coord.x+1, coord.y) {
				playersBuffer[Player{coord.x + 1, coord.y}] = true
			}

			// Down
			if isSafe(coord.x, coord.y+1) {
				playersBuffer[Player{coord.x, coord.y + 1}] = true
			}

			// Left
			if isSafe(coord.x-1, coord.y) {
				playersBuffer[Player{coord.x - 1, coord.y}] = true
			}

			// Up
			if isSafe(coord.x, coord.y-1) {
				if coord.y-1 == 0 {
					minute++
					break Outer2
				}
				playersBuffer[Player{coord.x, coord.y - 1}] = true
			}

			// Stay
			if isSafe(coord.x, coord.y) {
				playersBuffer[Player{coord.x, coord.y}] = true
			}
		}

		minute++
	}

	playersBuffer = make(map[Player]bool)
	playersBuffer[Player{1, 0}] = true
Outer3:
	for {
		// Move Blizzard
		for i := range blizzards {
			blizzards[i].Move()
		}

		// fmt.Println("Minute :", minute+1)
		// printMap()

		// Move Player
		players = playersBuffer
		playersBuffer = make(map[Player]bool)
		for coord := range players {
			// Right
			if isSafe(coord.x+1, coord.y) {
				playersBuffer[Player{coord.x + 1, coord.y}] = true
			}

			// Down
			if isSafe(coord.x, coord.y+1) {
				if coord.y+1 == len(grid)-1 {
					break Outer3
				}
				playersBuffer[Player{coord.x, coord.y + 1}] = true
			}

			// Left
			if isSafe(coord.x-1, coord.y) {
				playersBuffer[Player{coord.x - 1, coord.y}] = true
			}

			// Up
			if isSafe(coord.x, coord.y-1) {
				playersBuffer[Player{coord.x, coord.y - 1}] = true
			}

			// Stay
			if isSafe(coord.x, coord.y) {
				playersBuffer[Player{coord.x, coord.y}] = true
			}
		}

		minute++
	}

	fmt.Println(minute)

}
