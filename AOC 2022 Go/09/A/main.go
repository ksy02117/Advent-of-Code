package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	head, tail coord
	visited    map[coord]bool
)

type coord struct {
	x, y int
}

// when head moves up and if tail is not touching, then tail only can go to one spot
// same logic applies to other directions
func moveUp() {
	head.y++
	if head.y-tail.y >= 2 {
		tail.x = head.x
		tail.y = head.y - 1
		visited[tail] = true
	}
}

func moveDown() {
	head.y--
	if head.y-tail.y <= -2 {
		tail.x = head.x
		tail.y = head.y + 1
		visited[tail] = true
	}
}

func moveLeft() {
	head.x--
	if head.x-tail.x <= -2 {
		tail.x = head.x + 1
		tail.y = head.y
		visited[tail] = true
	}
}

func moveRight() {
	head.x++
	if head.x-tail.x >= 2 {
		tail.x = head.x - 1
		tail.y = head.y
		visited[tail] = true
	}
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	visited = make(map[coord]bool)
	visited[coord{0, 0}] = true
	fmt.Println("avc")

	var direction byte
	var amount int
	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%c %d", &direction, &amount)

		for i := 0; i < amount; i++ {
			switch direction {
			case 'U':
				moveUp()
			case 'D':
				moveDown()
			case 'R':
				moveRight()
			case 'L':
				moveLeft()
			}
		}
	}

	fmt.Println(len(visited))
}
