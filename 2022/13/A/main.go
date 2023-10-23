package main

import (
	"bufio"
	"os"
)

type Element struct {
	isNumber bool
	number   int
	list     []Element
}

func NewElement(input string, i *int) Element {
	var out Element

	for {
		if input[*i] == '[' {
			out.isNumber = false
			out.list = append(out.list, NewElement(input, i))
		}
		if input[*i] >= '0' && input[*i] < '9' {

		}
	}

	return out
}

func Compare(left, right *Element) int {
	if left.isNumber && right.isNumber {
		if left.number < right.number {
			return 1
		} else if left.number == right.number {
			return 0
		} else {
			return -1
		}
	}

	if left.isNumber && !right.isNumber {
		return Compare(&Element{false, 0, []Element{{true, left.number, nil}}}, right)
	}

	if !left.isNumber && right.isNumber {
		return Compare(left, &Element{false, 0, []Element{{true, right.number, nil}}})
	}

	if !left.isNumber && !right.isNumber {
		i := 0
		for ; i < len(left.list) && i < len(right.list); i++ {
			out := Compare(&left.list[i], &right.list[i])
			if out != 0 {
				return out
			}
		}

		if i == len(left.list) && i < len(right.list) {
			return 1
		}

		if i < len(left.list) && i == len(right.list) {
			return -1
		}

		return 0
	}

	panic("should not be reachable")
}

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	index, sum := 1, 0
	for sc.Scan() {

	}
}
