package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	value    byte
	position int
}

type stack struct {
	elements []node
}

func (s *stack) Process(value byte, position int) int {
	for len(s.elements) != 0 && s.elements[len(s.elements)-1].value < value {
		s.elements = s.elements[:len(s.elements)-1]
	}

	s.elements = append(s.elements, node{value, position})
	if len(s.elements) == 1 {
		return position
	}

	return position - s.elements[len(s.elements)-2].position
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var grid [][]byte
	var answer [][]int
	for sc.Scan() {
		grid = append(grid, []byte(sc.Text()))
		answer = append(answer, make([]int, len(grid[0])))

		for i := range answer[len(answer)-1] {
			answer[len(answer)-1][i] = 1
		}
	}

	height, width := len(grid), len(grid[0])
	maxRight, maxLeft := make([]stack, height), make([]stack, height)
	maxTop, maxBottom := make([]stack, width), make([]stack, width)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// Top
			answer[i][j] *= maxTop[j].Process(grid[i][j], i)

			// Bottom
			answer[height-i-1][j] *= maxBottom[j].Process(grid[height-i-1][j], i)

			// Left
			answer[i][j] *= maxLeft[i].Process(grid[i][j], j)

			// Right
			answer[i][width-j-1] *= maxRight[i].Process(grid[i][width-j-1], j)
		}
	}

	max := 0
	for i := range answer {
		for j := range answer[i] {
			if max < answer[i][j] {
				max = answer[i][j]
			}
		}
	}

	fmt.Println(max)
}
