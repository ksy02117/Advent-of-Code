package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	x, y, z int
}

var (
	grid  [][][]byte
	lavas []Coord

	maxX int
	maxY int
	maxZ int
)

func fillWaterDFS(c Coord) {
	if grid[c.x][c.y][c.z] != 0 {
		return
	}

	grid[c.x][c.y][c.z] = 'O'

	if c.x > 0 {
		fillWaterDFS(Coord{c.x - 1, c.y, c.z})
	}
	if c.x < maxX {
		fillWaterDFS(Coord{c.x + 1, c.y, c.z})
	}
	if c.y > 0 {
		fillWaterDFS(Coord{c.x, c.y - 1, c.z})
	}
	if c.y < maxY {
		fillWaterDFS(Coord{c.x, c.y + 1, c.z})
	}
	if c.z > 0 {
		fillWaterDFS(Coord{c.x, c.y, c.z - 1})
	}
	if c.z < maxZ {
		fillWaterDFS(Coord{c.x, c.y, c.z + 1})
	}
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var coord Coord
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%d,%d,%d", &coord.x, &coord.y, &coord.z)

		// transform to avoid 0 coordinates
		// if lava touches boundary we cannot dfs the water placements
		coord.x++
		coord.y++
		coord.z++

		if coord.x+1 > maxX {
			maxX = coord.x + 1
		}
		if coord.y+1 > maxY {
			maxY = coord.y + 1
		}
		if coord.z+1 > maxZ {
			maxZ = coord.z + 1
		}

		lavas = append(lavas, coord)
	}

	fmt.Println(maxX, maxY, maxZ)

	// initialize grid
	for x := 0; x <= maxX; x++ {
		plane := make([][]byte, 0)
		for y := 0; y <= maxY; y++ {
			row := make([]byte, maxZ+1)
			plane = append(plane, row)
		}
		grid = append(grid, plane)
	}

	// place lava
	for i := range lavas {
		grid[lavas[i].x][lavas[i].y][lavas[i].z] = '#'
		if lavas[i].x == 22 {
			fmt.Println(lavas[i])
		}
	}

	// fill water dfs from 0,0,0
	fillWaterDFS(Coord{0, 0, 0})

	for x := range grid {
		for y := range grid[x] {
			for z := range grid[x][y] {
				if grid[x][y][z] == 0 {
					fmt.Print(".")
				} else {
					fmt.Printf("%c", grid[x][y][z])
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	cnt := 0
	for i := range lavas {
		x, y, z := lavas[i].x, lavas[i].y, lavas[i].z

		if grid[x-1][y][z] == 'O' {
			cnt++
		}
		if grid[x+1][y][z] == 'O' {
			cnt++
		}
		if grid[x][y-1][z] == 'O' {
			cnt++
		}
		if grid[x][y+1][z] == 'O' {
			cnt++
		}
		if grid[x][y][z-1] == 'O' {
			cnt++
		}
		if grid[x][y][z+1] == 'O' {
			cnt++
		}
	}

	fmt.Println(cnt)
}
