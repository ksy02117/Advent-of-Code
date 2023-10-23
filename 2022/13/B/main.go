package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Element struct {
	isNumber bool
	number   int
	list     []Element
}

func NewElement(input string, i *int) Element {
	out := Element{false, 0, make([]Element, 0)}

	for {
		if input[*i] == '[' {
			*i++
			out.list = append(out.list, NewElement(input, i))
		}
		if input[*i] >= '0' && input[*i] <= '9' {
			num := 0
			for input[*i] >= '0' && input[*i] <= '9' {
				num *= 10
				num += int(input[*i] - '0')
				*i++
			}
			out.list = append(out.list, Element{true, num, nil})
		}
		if input[*i] == ']' {
			*i++
			return out
		}
		*i++
	}
}

func (e Element) String() string {
	if e.isNumber {
		return fmt.Sprintf("%d", e.number)
	}
	return fmt.Sprintf("%v", e.list)
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

	elements := make([]Element, 0)
	i := 1
	elements = append(elements, NewElement("[[2]]", &i))
	i = 1
	elements = append(elements, NewElement("[[6]]", &i))
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}

		i = 1
		elements = append(elements, NewElement(line, &i))
	}

	sort.SliceStable(elements, func(i, j int) bool {
		if Compare(&elements[i], &elements[j]) == 1 {
			return true
		} else {
			return false
		}
	})

	answer := 1
	for i := range elements {
		line := elements[i].String()
		if line == "[[2]]" {
			answer *= i + 1
		}
		if line == "[[6]]" {
			answer *= i + 1
		}
	}

	fmt.Println(answer)
}
