package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ksy02117/Advent-of-Code/2022/05/stack"
)

// used stack to lIFO the movement
func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)
	var stacks [9]stack.Stack

	for sc.Scan() {
		line := sc.Text()

		if line[1] == '1' {
			break
		}

		for i, j := 1, 0; i <= len(line); i, j = i+4, j+1 {
			if line[i] != ' ' {
				stacks[j].Push(line[i])
			}
		}
	}

	sc.Scan()
	for i := 0; i < len(stacks); i++ {
		stacks[i].Reverse()
	}

	var cnt, from, to int
	for sc.Scan() {
		line := sc.Text()

		fmt.Sscanf(line, "move %d from %d to %d", &cnt, &from, &to)

		for i := 0; i < cnt; i++ {
			stacks[to-1].Push(stacks[from-1].Pop())
		}
	}

	for _, stack := range stacks {
		fmt.Printf("%c", stack.Peek())
	}
}
