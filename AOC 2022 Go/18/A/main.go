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
	coords []Coord
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	cnt := 0
	for sc.Scan() {
		var coord Coord
		fmt.Sscanf(sc.Text(), "%d,%d,%d", &coord.x, &coord.y, &coord.z)

		// assume every side is open and subtract it later
		cnt += 6
		for i := range coords {
			if coord.x == coords[i].x && coord.y == coords[i].y && coord.z == coords[i].z+1 {
				cnt -= 2
				continue
			}
			if coord.x == coords[i].x && coord.y == coords[i].y && coord.z == coords[i].z-1 {
				cnt -= 2
				continue
			}
			if coord.x == coords[i].x && coord.y == coords[i].y+1 && coord.z == coords[i].z {
				cnt -= 2
				continue
			}
			if coord.x == coords[i].x && coord.y == coords[i].y-1 && coord.z == coords[i].z {
				cnt -= 2
				continue
			}
			if coord.x == coords[i].x+1 && coord.y == coords[i].y && coord.z == coords[i].z {
				cnt -= 2
				continue
			}
			if coord.x == coords[i].x-1 && coord.y == coords[i].y && coord.z == coords[i].z {
				cnt -= 2
				continue
			}
		}

		coords = append(coords, coord)
	}

	fmt.Println(cnt)
}
