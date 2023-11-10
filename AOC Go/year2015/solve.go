package year2015

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ksy02117/Advent-of-Code/AOC-Go/year2015/day01"
)

var (
	sols = map[string]func(string) (string, error){
		"1A": day01.SolveA,
	}
)

func Solve(day int, part string) (string, error) {
	k := strconv.Itoa(day) + part
	sol, ok := sols[k]
	if !ok {
		return "", errors.New("Solution Not Implemented")
	}

	filePath := fmt.Sprintf("year2015/day%02d/input.txt", day)
	result, err := sol(filePath)
	return result, err
}
