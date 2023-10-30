package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	sum := 0
	for sc.Scan() {
		line := sc.Text()
		point := calPoint2(line)
		fmt.Printf("[%v]:%v\n", line, point)
		sum += point
	}

	fmt.Println(sum)
}

// the point function for A
func calPoint1(line string) int {
	// A, X Rock
	// B, Y Paper
	// C, Z Scissors
	switch line {
	case "A X":
		return 1 + 3
	case "A Y":
		return 2 + 6
	case "A Z":
		return 3 + 0
	case "B X":
		return 1 + 0
	case "B Y":
		return 2 + 3
	case "B Z":
		return 3 + 6
	case "C X":
		return 1 + 6
	case "C Y":
		return 2 + 0
	case "C Z":
		return 3 + 3
	}
	panic("no case")
}

// the point function for B
func calPoint2(line string) int {
	// A Rock
	// B Paper
	// C Scissors
	// X loss
	// Y Draw
	// Z Win
	switch line {
	case "A X":
		return 3 + 0
	case "A Y":
		return 1 + 3
	case "A Z":
		return 2 + 6
	case "B X":
		return 1 + 0
	case "B Y":
		return 2 + 3
	case "B Z":
		return 3 + 6
	case "C X":
		return 2 + 0
	case "C Y":
		return 3 + 3
	case "C Z":
		return 1 + 6
	}
	panic("no case")
}
