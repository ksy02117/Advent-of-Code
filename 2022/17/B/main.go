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
		if grid[r.y-1][r.x] || grid[r.y-1][r.x+1] || grid[r.y-1][r.x+2] || grid[r.y-1][r.x+3] {
			return false
		}
	case 1:
		if grid[r.y][r.x] || grid[r.y-1][r.x+1] || grid[r.y][r.x+2] {
			return false
		}
	case 2:
		if grid[r.y-1][r.x] || grid[r.y-1][r.x+1] || grid[r.y-1][r.x+2] {
			return false
		}
	case 3:
		if grid[r.y-1][r.x] {
			return false
		}
	case 4:
		if grid[r.y-1][r.x] || grid[r.y-1][r.x+1] {
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
		if grid[r.y][r.x-1] {
			return false
		}
	case 1:
		if grid[r.y][r.x] || grid[r.y+1][r.x-1] || grid[r.y+2][r.x] {
			return false
		}
	case 2:
		if grid[r.y][r.x-1] || grid[r.y+1][r.x+1] || grid[r.y+2][r.x+1] {
			return false
		}
	case 3:
		if grid[r.y][r.x-1] || grid[r.y+1][r.x-1] || grid[r.y+2][r.x-1] || grid[r.y+3][r.x-1] {
			return false
		}
	case 4:
		if grid[r.y][r.x-1] || grid[r.y+1][r.x-1] {
			return false
		}
	}

	r.x--
	return true
}

func (r *Rock) MoveRight() bool {
	switch r.ID % 5 {
	case 0:
		if r.x+4 >= 7 || grid[r.y][r.x+4] {
			return false
		}
	case 1:
		if r.x+3 >= 7 || grid[r.y][r.x+2] || grid[r.y+1][r.x+3] || grid[r.y+2][r.x+2] {
			return false
		}
	case 2:
		if r.x+3 >= 7 || grid[r.y][r.x+3] || grid[r.y+1][r.x+3] || grid[r.y+2][r.x+3] {
			return false
		}
	case 3:
		if r.x+1 >= 7 || grid[r.y][r.x+1] || grid[r.y+1][r.x+1] || grid[r.y+2][r.x+1] || grid[r.y+3][r.x+1] {
			return false
		}
	case 4:
		if r.x+2 >= 7 || grid[r.y][r.x+2] || grid[r.y+1][r.x+2] {
			return false
		}
	}

	r.x++
	return true
}

func (r *Rock) PlaceDown() (int, int) {
	switch r.ID % 5 {
	case 0:
		grid[r.y][r.x] = true
		grid[r.y][r.x+1] = true
		grid[r.y][r.x+2] = true
		grid[r.y][r.x+3] = true
		return r.y, r.y
	case 1:
		grid[r.y+2][r.x+1] = true
		grid[r.y+1][r.x] = true
		grid[r.y+1][r.x+1] = true
		grid[r.y+1][r.x+2] = true
		grid[r.y][r.x+1] = true
		return r.y, r.y + 2
	case 2:
		grid[r.y+2][r.x+2] = true
		grid[r.y+1][r.x+2] = true
		grid[r.y][r.x] = true
		grid[r.y][r.x+1] = true
		grid[r.y][r.x+2] = true
		return r.y, r.y + 2
	case 3:
		grid[r.y+3][r.x] = true
		grid[r.y+2][r.x] = true
		grid[r.y+1][r.x] = true
		grid[r.y][r.x] = true
		return r.y, r.y + 3
	case 4:
		grid[r.y+1][r.x] = true
		grid[r.y+1][r.x+1] = true
		grid[r.y][r.x] = true
		grid[r.y][r.x+1] = true
		return r.y, r.y + 1
	}

	panic("PlaceDown:: Should not be reachable")
}

type State struct {
	grid          [][]bool
	movementID    int
	rockID        int
	removedHeight int
}

var (
	grid          [][]bool
	movement      string
	movementID    int = 0
	maxHeight     int = -1
	removedHeight int = 0

	states []State
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	sc.Scan()
	movement = sc.Text()

	// for loop for rock creation
	const MAX_ROCKS = 1000000000000
	for rockID := 0; rockID < MAX_ROCKS; rockID++ {
		if rockID%1000000000 == 0 {
			fmt.Println(rockID)
		}
		// increase map size
		for len(grid) <= maxHeight+8 {
			grid = append(grid, make([]bool, 7))
		}

		rock := Rock{2, maxHeight + 4, rockID}

		// for loop foreach movement
		for {
			switch movement[movementID%(len(movement))] {
			case '<':
				// fmt.Println(rock, "Moving Left")
				rock.MoveLeft()
			case '>':
				// fmt.Println(rock, "Moving Right")
				rock.MoveRight()
			}
			movementID = (movementID + 1) % len(movement)

			if rock.MoveDown() {
				continue
			}

			height1, height2 := rock.PlaceDown()

			// update max height
			if maxHeight < height2 {
				maxHeight = height2
			}

			// closing detection
			for y := height2; y >= height1; y-- {
				isFilled := true
				for x := range grid[y] {
					if !grid[y][x] {
						isFilled = false
					}
				}

				if isFilled {
					// remove unaccessible rows
					fmt.Println("Filled a row")
					removedHeight += y + 1
					maxHeight -= y + 1
					grid = grid[y+1:]

					// state save
					var state State
					state.grid = make([][]bool, 0)
					state.rockID = rockID
					state.movementID = movementID
					state.removedHeight = removedHeight
					for y := 0; y < maxHeight; y++ {
						state.grid = append(state.grid, make([]bool, 7))
						for x := 0; x < 7; x++ {
							state.grid[y][x] = grid[y][x]
						}
					}

					// check for identical state
					for i := range states {
						if states[i].movementID != state.movementID {
							continue
						}
						if states[i].rockID%5 != state.rockID%5 {
							continue
						}
						if len(states[i].grid) != len(state.grid) {
							continue
						}
						isGridSame := true
						for y := range state.grid {
							for x := range state.grid[y] {
								if states[i].grid[y][x] != state.grid[y][x] {
									isGridSame = false
								}
							}
						}
						if !isGridSame {
							continue
						}

						fmt.Println("identical state found")
						rockIDDif := state.rockID - states[i].rockID
						removedHeightDif := state.removedHeight - states[i].removedHeight
						for rockID+rockIDDif < MAX_ROCKS {
							rockID += rockIDDif
							removedHeight += removedHeightDif
						}

						break
					}
					states = append(states, state)
					break
				}
			}

			break
		}
	}

	// for y := maxHeight + 2; y >= 0; y-- {
	// 	fmt.Printf("%4d : ", y)
	// 	for x := range grid[y] {
	// 		if grid[y][x] {
	// 			fmt.Print("#")
	// 		} else {
	// 			fmt.Print(".")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println(removedHeight + maxHeight + 1)
}
