package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	snake   [10]coord
	visited map[coord]bool
)

type coord struct {
	x, y int
}

// by removing the case where next node is touching previous node
// any difference in x and y will move tail in that direction
func (tail *coord) follow(head coord) {
	vec := coord{head.x - tail.x, head.y - tail.y}

	if vec.x < 2 && vec.x > -2 && vec.y < 2 && vec.y > -2 {
		return
	}

	if vec.x > 0 {
		tail.x++
	}
	if vec.x < 0 {
		tail.x--
	}
	if vec.y > 0 {
		tail.y++
	}
	if vec.y < 0 {
		tail.y--
	}
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	visited = make(map[coord]bool)

	var direction byte
	var amount int
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%c %d", &direction, &amount)

		for i := 0; i < amount; i++ {
			switch direction {
			case 'U':
				snake[0].y++
			case 'D':
				snake[0].y--
			case 'R':
				snake[0].x++
			case 'L':
				snake[0].x--
			}

			for j := 1; j < 10; j++ {
				snake[j].follow(snake[j-1])
			}

			visited[snake[9]] = true
		}
	}

	fmt.Println(len(visited))
}
