package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	var grid [][]byte
	var visible [][]bool
	for sc.Scan() {
		grid = append(grid, []byte(sc.Text()))
		visible = append(visible, make([]bool, len(grid[0])))
	}

	height, width := len(grid), len(grid[0])
	maxRight, maxLeft := make([]byte, height), make([]byte, height)
	maxTop, maxBottom := make([]byte, width), make([]byte, width)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// Top
			if grid[i][j] > maxTop[j] {
				visible[i][j] = true
				maxTop[j] = grid[i][j]
			}

			// Bottom
			if grid[height-i-1][j] > maxBottom[j] {
				visible[height-i-1][j] = true
				maxBottom[j] = grid[height-i-1][j]
			}

			// Left
			if grid[i][j] > maxLeft[i] {
				visible[i][j] = true
				maxLeft[i] = grid[i][j]
			}

			// Right
			if grid[i][width-j-1] > maxRight[i] {
				visible[i][width-j-1] = true
				maxRight[i] = grid[i][width-j-1]
			}
		}
	}

	cnt := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if visible[i][j] {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
