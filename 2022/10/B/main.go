package main

import (
	"bufio"
	"fmt"
	"os"
)

type instruction struct {
	finishTime, value int
}

var (
	grid [6][40]byte
	x    int         = 1
	ci   instruction = instruction{0, 0}
)

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for cycle := 0; ; cycle++ {
		if cycle/40 >= 6 {
			break
		}

		if cycle == ci.finishTime {
			x += ci.value
			read := sc.Scan()
			if !read {
				break
			}
			line := sc.Text()

			if line == "noop" {
				ci.finishTime = cycle + 1
				ci.value = 0
			} else {
				fmt.Sscanf(line, "addx %d", &ci.value)
				ci.finishTime = cycle + 2
			}
		}

		if cycle%40-x <= 1 && cycle%40-x >= -1 {
			grid[cycle/40][cycle%40] = '#'
		} else {
			grid[cycle/40][cycle%40] = '.'
		}
	}
	printGrid()
}

func printGrid() {
	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
}
