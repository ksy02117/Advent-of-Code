package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	x, y int
}

type Elf struct {
	loc Coord
}

func isElf(x, y int) bool {
	for i := range elves {
		if elves[i].loc.x == x && elves[i].loc.y == y {
			return true
		}
	}
	return false
}

func (e *Elf) Propose() (Coord, bool) {
	if !isElf(e.loc.x-1, e.loc.y-1) && !isElf(e.loc.x, e.loc.y-1) && !isElf(e.loc.x+1, e.loc.y-1) &&
		!isElf(e.loc.x-1, e.loc.y) && !isElf(e.loc.x+1, e.loc.y) &&
		!isElf(e.loc.x-1, e.loc.y+1) && !isElf(e.loc.x, e.loc.y+1) && !isElf(e.loc.x+1, e.loc.y+1) {
		return Coord{}, false
	}

	for i := 0; i < 4; i++ {
		switch (dir + i) % 4 {
		case 0: // North (UP)
			if !isElf(e.loc.x-1, e.loc.y-1) && !isElf(e.loc.x, e.loc.y-1) && !isElf(e.loc.x+1, e.loc.y-1) {
				return Coord{e.loc.x, e.loc.y - 1}, true
			}
		case 1: // South (DOWN)
			if !isElf(e.loc.x-1, e.loc.y+1) && !isElf(e.loc.x, e.loc.y+1) && !isElf(e.loc.x+1, e.loc.y+1) {
				return Coord{e.loc.x, e.loc.y + 1}, true
			}
		case 2: // West (LEFT)
			if !isElf(e.loc.x-1, e.loc.y-1) && !isElf(e.loc.x-1, e.loc.y) && !isElf(e.loc.x-1, e.loc.y+1) {
				return Coord{e.loc.x - 1, e.loc.y}, true
			}
		case 3: // East (RIGHT)
			if !isElf(e.loc.x+1, e.loc.y-1) && !isElf(e.loc.x+1, e.loc.y) && !isElf(e.loc.x+1, e.loc.y+1) {
				return Coord{e.loc.x + 1, e.loc.y}, true
			}
		}
	}
	return Coord{}, false
}

func printMap() {
	minX, maxX, minY, maxY := 0, 0, 0, 0

	for i := range elves {
		if minX > elves[i].loc.x {
			minX = elves[i].loc.x
		} else if maxX < elves[i].loc.x {
			maxX = elves[i].loc.x
		}
		if minY > elves[i].loc.y {
			minY = elves[i].loc.y
		} else if maxY < elves[i].loc.y {
			maxY = elves[i].loc.y
		}
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if isElf(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

var (
	elves        []Elf = make([]Elf, 0)
	destinations map[Coord][]*Elf
	dir          int
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for y := 0; sc.Scan(); y++ {
		line := sc.Text()

		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				elves = append(elves, Elf{loc: Coord{x, y}})
			}
		}
	}

	// fmt.Println("Round", 0)
	// printMap()

	round := 1
	for {
		// Propose
		destinations = make(map[Coord][]*Elf)
		for j := range elves {
			dest, k := elves[j].Propose()
			if !k {
				continue
			}

			if _, k = destinations[dest]; !k {
				destinations[dest] = make([]*Elf, 0)
			}

			destinations[dest] = append(destinations[dest], &elves[j])
		}

		// Move
		cnt := 0
		for coord := range destinations {
			if len(destinations[coord]) == 1 {
				destinations[coord][0].loc = coord
				cnt++
			}
		}

		if cnt == 0 {
			break
		}

		dir = (dir + 1) % 4
		round++

		// fmt.Println("Round", i+1)
		// printMap()
		// fmt.Println()
	}

	fmt.Println(round)
}
