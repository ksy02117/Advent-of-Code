package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	number int
	order  int
}

type List struct {
	elements []Pair
}

func (l List) String() string {
	out := "["
	for i := range l.elements {
		out += fmt.Sprint(l.elements[i].number) + " "
	}
	out += "]"
	return out
}

func (l *List) ShiftLeft(idx int) int {
	if idx == 0 {
		l.elements = append(l.elements[1:], l.elements[0])
		idx = len(l.elements) - 1
	}

	tmp := l.elements[idx]
	l.elements[idx] = l.elements[idx-1]
	l.elements[idx-1] = tmp

	return idx - 1
}

func (l *List) ShiftRight(idx int) int {
	if idx >= len(l.elements)-1 {
		l.elements = append(l.elements[len(l.elements)-1:], l.elements[:len(l.elements)-1]...)
		idx = 0
	}

	tmp := l.elements[idx]
	l.elements[idx] = l.elements[idx+1]
	l.elements[idx+1] = tmp

	return idx + 1
}

var list List

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	number := 0
	for i := 0; sc.Scan(); i++ {
		fmt.Sscanf(sc.Text(), "%d", &number)
		list.elements = append(list.elements, Pair{number, i})
	}

	fmt.Println(list)

	for i := 0; i < len(list.elements); i++ {
		var number, idx int
		for idx = 0; idx < len(list.elements); idx++ {
			if list.elements[idx].order == i {
				number = list.elements[idx].number
				break
			}
		}

		for number != 0 {
			if number > 0 {
				idx = list.ShiftRight(idx)
				number--
			} else {
				idx = list.ShiftLeft(idx)
				number++
			}
		}
	}

	var idx int
	for i := range list.elements {
		if list.elements[i].number == 0 {
			idx = i
			break
		}
	}

	sum := 0
	for i := 1; i <= 3000; i++ {
		idx = (idx + 1) % len(list.elements)

		if i == 1000 || i == 2000 || i == 3000 {
			fmt.Println(i, list.elements[idx].number)
			sum += list.elements[idx].number
		}
	}
	fmt.Println(sum)
}
