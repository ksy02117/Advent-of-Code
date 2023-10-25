package main

import (
	"bufio"
	"fmt"
	"os"
)

type Rock struct {
	x, y, ID int
}

func (r *Rock) MoveDown() bool {
	if r.y == 0 {
		return false
	}
	switch r.ID % 5 {
	case 0:
		if grid[r.y-1][r.x] != 0 || grid[r.y-1][r.x+1] != 0 || grid[r.y-1][r.x+2] != 0 || grid[r.y-1][r.x+3] != 0 {
			return false
		}
	case 1:
		if grid[r.y][r.x] != 0 || grid[r.y-1][r.x+1] != 0 || grid[r.y][r.x+2] != 0 {
			return false
		}
	case 2:
		if grid[r.y-1][r.x] != 0 || grid[r.y-1][r.x+1] != 0 || grid[r.y-1][r.x+2] != 0 {
			return false
		}
	case 3:
		if grid[r.y-1][r.x] != 0 {
			return false
		}
	case 4:
		if grid[r.y-1][r.x] != 0 || grid[r.y-1][r.x+1] != 0 {
			return false
		}
	}

	r.y--
	return true
}

func (r *Rock) MoveLeft() bool {
	if r.x == 0 {
		return false
	}
	switch r.ID % 5 {
	case 0:
		if grid[r.y][r.x-1] != 0 {
			return false
		}
		r.x--
		return true
	case 1:
		if grid[r.y][r.x] != 0 || grid[r.y+1][r.x-1] != 0 || grid[r.y+2][r.x] != 0 {
			return false
		}
		r.x--
		return true
	case 2:
		if grid[r.y][r.x-1] != 0 || grid[r.y+1][r.x+1] != 0 || grid[r.y+2][r.x+1] != 0 {
			return false
		}
		r.x--
		return true
	case 3:
		if grid[r.y][r.x-1] != 0 || grid[r.y+1][r.x-1] != 0 || grid[r.y+2][r.x-1] != 0 || grid[r.y+3][r.x-1] != 0 {
			return false
		}
		r.x--
		return true
	case 4:
		if grid[r.y][r.x-1] != 0 || grid[r.y+1][r.x-1] != 0 {
			return false
		}
		r.x--
		return true
	}

	panic("MoveLeft:: should not be reachable")
}

func (r *Rock) MoveRight() bool {
	switch r.ID % 5 {
	case 0:
		if r.x+4 >= 7 || grid[r.y][r.x+4] != 0 {
			return false
		}
		r.x++
		return true
	case 1:
		if r.x+3 >= 7 || grid[r.y][r.x+2] != 0 || grid[r.y+1][r.x+3] != 0 || grid[r.y+2][r.x+2] != 0 {
			return false
		}
		r.x++
		return true
	case 2:
		if r.x+3 >= 7 || grid[r.y][r.x+3] != 0 || grid[r.y+1][r.x+3] != 0 || grid[r.y+2][r.x+3] != 0 {
			return false
		}
		r.x++
		return true
	case 3:
		if r.x+1 >= 7 || grid[r.y][r.x+1] != 0 || grid[r.y+1][r.x+1] != 0 || grid[r.y+2][r.x+1] != 0 || grid[r.y+3][r.x+1] != 0 {
			return false
		}
		r.x++
		return true
	case 4:
		if r.x+2 >= 7 || grid[r.y][r.x+2] != 0 || grid[r.y+1][r.x+2] != 0 {
			return false
		}
		r.x++
		return true
	}

	panic("MoveRight:: should not be reachable")
}

func (r *Rock) PlaceDown() int {
	switch r.ID % 5 {
	case 0:
		grid[r.y][r.x] = '-'
		grid[r.y][r.x+1] = '-'
		grid[r.y][r.x+2] = '-'
		grid[r.y][r.x+3] = '-'
		return r.y
	case 1:
		grid[r.y+2][r.x+1] = '+'
		grid[r.y+1][r.x] = '+'
		grid[r.y+1][r.x+1] = '+'
		grid[r.y+1][r.x+2] = '+'
		grid[r.y][r.x+1] = '+'
		return r.y + 2
	case 2:
		grid[r.y+2][r.x+2] = 'J'
		grid[r.y+1][r.x+2] = 'J'
		grid[r.y][r.x] = 'J'
		grid[r.y][r.x+1] = 'J'
		grid[r.y][r.x+2] = 'J'
		return r.y + 2
	case 3:
		grid[r.y+3][r.x] = 'I'
		grid[r.y+2][r.x] = 'I'
		grid[r.y+1][r.x] = 'I'
		grid[r.y][r.x] = 'I'
		return r.y + 3
	case 4:
		grid[r.y+1][r.x] = 'O'
		grid[r.y+1][r.x+1] = 'O'
		grid[r.y][r.x] = 'O'
		grid[r.y][r.x+1] = 'O'
		return r.y + 1
	}

	panic("PlaceDown:: Should not be reachable")
}

var (
	grid       [][]byte
	movement   string
	movementID int = 0
	height     int = -1
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	sc.Scan()
	movement = sc.Text()

	// for loop for rock creation
	for rockID := 0; rockID < 2022; rockID++ {
		// increase map size
		for len(grid) <= height+8 {
			grid = append(grid, make([]byte, 7))
		}

		rock := Rock{2, height + 4, rockID}

		// for loop foreach movement
		for {
			if movementID < 0 {
				panic("movement Wrap Around")
			}
			switch movement[movementID%(len(movement))] {
			case '<':
				// fmt.Println(rock, "Moving Left")
				rock.MoveLeft()
			case '>':
				// fmt.Println(rock, "Moving Right")
				rock.MoveRight()
			}
			movementID++

			if rock.MoveDown() {
				continue
			}

			newHeight := rock.PlaceDown()
			if newHeight > height {
				height = newHeight
			}
			break
		}
	}

	// for y := height + 2; y >= height-20; y-- {
	// 	fmt.Printf("%4d : ", y)
	// 	for x := range grid[y] {
	// 		if grid[y][x] == 0 {
	// 			fmt.Print(".")
	// 		} else {
	// 			fmt.Printf("%c", grid[y][x])
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println(height + 1)
}
