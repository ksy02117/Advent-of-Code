package main

// Not a General Solution
// To implement Generic Solution generate automatic ways to make edge map

import (
	"bufio"
	"fmt"
	"os"
)

type Plane struct {
	grid [52][52]byte
}

type Player struct {
	plane  int
	x, y   int
	facing int
}

func (p *Player) Move() bool {
	var destination Coord

	switch p.facing {
	case 0: // Right
		switch face[p.plane].grid[p.y][p.x+1] {
		case '#': // Stop
			return false
		case '.': // Continue
			p.x++
			return true
		case ' ':
			destination = edge[Coord{face: p.plane, x: p.x + 1, y: p.y}]
		}
	case 1: // Down
		switch face[p.plane].grid[p.y+1][p.x] {
		case '#': // Stop
			return false
		case '.': // Continue
			p.y++
			return true
		case ' ':
			destination = edge[Coord{face: p.plane, x: p.x, y: p.y + 1}]
		}
	case 2: // Left
		switch face[p.plane].grid[p.y][p.x-1] {
		case '#': // Stop
			return false
		case '.': // Continue
			p.x--
			return true
		case ' ':
			destination = edge[Coord{face: p.plane, x: p.x - 1, y: p.y}]
		}
	case 3: // Up
		switch face[p.plane].grid[p.y-1][p.x] {
		case '#': // Stop
			return false
		case '.': // Continue
			p.y--
			return true
		case ' ':
			destination = edge[Coord{face: p.plane, x: p.x, y: p.y - 1}]
		}
	}

	// Wrap Around
	fmt.Printf("mapping around from %d{%d, %d} to %d{%d, %d}\n", p.plane, p.x, p.y, destination.face, destination.x, destination.y)
	if destination.x == 0 {
		if face[destination.face].grid[destination.y][destination.x+1] == '#' {
			return false
		} else {
			p.facing = 0 // right
			p.plane = destination.face
			p.x = destination.x + 1
			p.y = destination.y
		}
	}
	if destination.y == 0 {
		if face[destination.face].grid[destination.y+1][destination.x] == '#' {
			return false
		} else {
			p.facing = 1 // down
			p.plane = destination.face
			p.x = destination.x
			p.y = destination.y + 1
		}
	}
	if destination.x == 51 {
		if face[destination.face].grid[destination.y][destination.x-1] == '#' {
			return false
		} else {
			p.facing = 2 // left
			p.plane = destination.face
			p.x = destination.x - 1
			p.y = destination.y
		}
	}
	if destination.y == 51 {
		if face[destination.face].grid[destination.y-1][destination.x] == '#' {
			return false
		} else {
			p.facing = 3 // up
			p.plane = destination.face
			p.x = destination.x
			p.y = destination.y - 1
		}
	}

	return true
}

func (p *Player) Turn(rotation int) {
	p.facing = (p.facing + 4 + rotation) % 4
}

func printMap() {
	for i := 0; i < 6; i++ {
		fmt.Println("face: ", i)
		for y := 0; y < len(face[i].grid); y++ {
			for x := 0; x < len(face[i].grid[y]); x++ {
				if face[i].grid[y][x] == 0 {
					fmt.Print("[X]")
				} else {
					fmt.Printf("[%c]", face[i].grid[y][x])
				}
			}
			fmt.Println()
		}
	}
}

type Coord struct {
	face, x, y int
}

var (
	fullGrid [][]byte = make([][]byte, 0)
	cube     [4][4]*Plane
	face     [6]*Plane
	edge     map[Coord]Coord
	player   Player
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	// fill full grid from map
	for sc.Scan() {
		if sc.Text() == "" {
			break
		}
		fullGrid = append(fullGrid, []byte(sc.Text()))
	}

	// get cube
	faceIdx := 0
	for y := 0; y < len(fullGrid); y++ {
		for x := 0; x < len(fullGrid[y]); x++ {
			if x%50 == 0 {
				if fullGrid[y][x] == ' ' {
					x += 49
					continue
				} else if y%50 == 0 {
					cube[y/50][x/50] = &Plane{}
					face[faceIdx] = cube[y/50][x/50]
					faceIdx++
				}
			}

			cube[y/50][x/50].grid[y%50+1][x%50+1] = fullGrid[y][x]
		}
	}

	// surround faces with ' '
	// "Same" square will be mapped to each other so the player travel to next square
	for i := 0; i < 6; i++ {
		for j := 1; j <= 50; j++ {
			face[i].grid[0][j] = ' '
			face[i].grid[51][j] = ' '
			face[i].grid[j][0] = ' '
			face[i].grid[j][51] = ' '
		}
	}

	// Set Edge map (not general solution)
	edge = make(map[Coord]Coord)
	for i := 1; i <= 50; i++ {
		// face 0 & 1
		edge[Coord{face: 0, x: 51, y: i}] = Coord{face: 1, x: 0, y: i}
		edge[Coord{face: 1, x: 0, y: i}] = Coord{face: 0, x: 51, y: i}

		// face 0 & 2
		edge[Coord{face: 0, x: i, y: 51}] = Coord{face: 2, x: i, y: 0}
		edge[Coord{face: 2, x: i, y: 0}] = Coord{face: 0, x: i, y: 51}

		// face 0 & 3
		edge[Coord{face: 0, x: 0, y: i}] = Coord{face: 3, x: 0, y: 51 - i}
		edge[Coord{face: 3, x: 0, y: i}] = Coord{face: 0, x: 0, y: 51 - i}

		// face 0 & 5
		edge[Coord{face: 0, x: i, y: 0}] = Coord{face: 5, x: 0, y: i}
		edge[Coord{face: 5, x: 0, y: i}] = Coord{face: 0, x: i, y: 0}

		// face 1 & 2
		edge[Coord{face: 1, x: i, y: 51}] = Coord{face: 2, x: 51, y: i}
		edge[Coord{face: 2, x: 51, y: i}] = Coord{face: 1, x: i, y: 51}

		// face 1 & 4
		edge[Coord{face: 1, x: 51, y: i}] = Coord{face: 4, x: 51, y: 51 - i}
		edge[Coord{face: 4, x: 51, y: i}] = Coord{face: 1, x: 51, y: 51 - i}

		// face 1 & 5
		edge[Coord{face: 1, x: i, y: 0}] = Coord{face: 5, x: i, y: 51}
		edge[Coord{face: 5, x: i, y: 51}] = Coord{face: 1, x: i, y: 0}

		// face 2 & 3
		edge[Coord{face: 2, x: 0, y: i}] = Coord{face: 3, x: i, y: 0}
		edge[Coord{face: 3, x: i, y: 0}] = Coord{face: 2, x: 0, y: i}

		// face 2 & 4
		edge[Coord{face: 2, x: i, y: 51}] = Coord{face: 4, x: i, y: 0}
		edge[Coord{face: 4, x: i, y: 0}] = Coord{face: 2, x: i, y: 51}

		// face 3 & 4
		edge[Coord{face: 3, x: 51, y: i}] = Coord{face: 4, x: 0, y: i}
		edge[Coord{face: 4, x: 0, y: i}] = Coord{face: 3, x: 51, y: i}

		// face 3 & 5
		edge[Coord{face: 3, x: i, y: 51}] = Coord{face: 5, x: i, y: 0}
		edge[Coord{face: 5, x: i, y: 0}] = Coord{face: 3, x: i, y: 51}

		// face 4 & 5
		edge[Coord{face: 4, x: i, y: 51}] = Coord{face: 5, x: 51, y: i}
		edge[Coord{face: 5, x: 51, y: i}] = Coord{face: 4, x: i, y: 51}

	}

	// set player location
	player.y = 1
	for x := 1; x <= 50; x++ {
		if face[0].grid[1][x] == '.' {
			player.x = x
			break
		}
	}

	fmt.Println(player)
	printMap()

	sc.Scan()
	var number int
	var rotation byte
	instruction := sc.Text()
	fmt.Println(instruction)

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

	fmt.Println(player.plane, player.y, player.x, player.facing)
	fmt.Println(1000*(player.y+100) + 4*player.x + player.facing)
}
