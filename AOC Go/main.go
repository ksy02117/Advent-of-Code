package main

import (
	"fmt"

	"github.com/ksy02117/Advent-of-Code/AOC-Go/year2015"
)

func main() {
	result, err := year2015.Solve(1, "A")
	if err == nil {
		fmt.Println(result)
	}
}
