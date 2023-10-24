package main

import (
	"bufio"
	"fmt"
	"os"
)

// loops text.
func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var v int
	x, cycle, sum := 1, 0, 0

	for sc.Scan() {
		line := sc.Text()

		if line == "noop" {
			cycle++
			if cycle%40 == 20 {
				fmt.Printf("cycle: %v, x: %v\n", cycle, x)
				sum += cycle * x
			}
			continue
		}

		// since addx takes two cycle we need to check for 20 and 21 case
		cycle += 2
		if cycle%40 == 20 || cycle%40 == 21 {
			fmt.Printf("cycle: %v, x: %v\n", (cycle/40*40 + 20), x)
			sum += (cycle/40*40 + 20) * x
		}
		fmt.Sscanf(line, "addx %d", &v)
		x += v
	}

	fmt.Println(sum)
}
