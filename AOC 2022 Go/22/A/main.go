package main

import (
	"bufio"
	"fmt"
	"os"
)

type Player struct {
	x, y   int
	facing int
}

func (p *Player) Move() bool {
	switch p.facing {
	case 0: // Right
		switch grid[p.y][p.x+1] {
		case '#': // Stop
			return false
		case '.': // Continue
			p.x++
		case ' ': // Wrap around
			x := p.x
			for grid[p.y][x-1] != ' ' {
				x--
			}
			if grid[p.y][x] == '.' {
				p.x = x
			}
		default:
			panic("")
		}
	case 1: // Down
		switch grid[p.y+1][p.x] {
		case '#': // Stop
			return false
		case '.': // Continue
			p.y++
		case ' ': // Wrap around
			y := p.y
			for grid[y-1][p.x] != ' ' {
				y--
			}
			if grid[y][p.x] == '.' {
				p.y = y
			}
		default:
			panic("")
		}
	case 2: // Left
		switch grid[p.y][p.x-1] {
		case '#': // Stop
			return false
		case '.': // Continue
			p.x--
		case ' ': // Wrap around
			x := p.x
			for grid[p.y][x+1] != ' ' {
				x++
			}
			if grid[p.y][x] == '.' {
				p.x = x
			}
		default:
			panic("")
		}
	case 3: // Top
		switch grid[p.y-1][p.x] {
		case '#': // Stop
			return false
		case '.': // Continue
			p.y--
		case ' ': // Wrap around
			y := p.y
			for grid[y+1][p.x] != ' ' {
				y++
			}
			if grid[y][p.x] == '.' {
				p.y = y
			}
		default:
			panic("")
		}
	default:
		panic("")
	}

	return true
}

func (p *Player) Turn(rotation int) {
	p.facing = (p.facing + 4 + rotation) % 4
}

func printGrid() {
	for y := range grid {
		for x := range grid[y] {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Println()
	}
}

var (
	grid   [][]byte = make([][]byte, 0)
	player Player

	maxWidth int = 0
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	// fill board with bordering ' ' so the other algorithms do not have to check for bounds
	grid = append(grid, make([]byte, 0))
	for sc.Scan() {
		line := sc.Text()

		if len(line) == 0 {
			break
		}

		row := make([]byte, 0)
		row = append(row, ' ')
		row = append(row, []byte(line)...)
		row = append(row, ' ')

		grid = append(grid, row)

		if maxWidth < len(row) {
			maxWidth = len(row)
		}
	}
	grid = append(grid, make([]byte, 0))

	for y := range grid {
		for len(grid[y]) != maxWidth {
			grid[y] = append(grid[y], ' ')
		}
	}

	// set player location
	player.y = 1
	for i := range grid[1] {
		if grid[1][i] == '.' {
			player.x = i
			break
		}
	}

	fmt.Println(player)
	printGrid()

	sc.Scan()
	var number int
	var rotation byte
	instruction := sc.Text()
	for {
		_, err := fmt.Sscanf(instruction, "%d%c%s", &number, &rotation, &instruction)

		// fmt.Println(number, string(rotation))

		for i := 0; i < number; i++ {
			if !player.Move() {
				break
			}
		}

		if err != nil {
			break
		}

		if rotation == 'R' {
			player.Turn(1)
		} else if rotation == 'L' {
			player.Turn(-1)
		} else {
			panic("")
		}
	}

	fmt.Println(player.y, player.x, player.facing)
	fmt.Println(1000*player.y + 4*player.x + player.facing)
}
