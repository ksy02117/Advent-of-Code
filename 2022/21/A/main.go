package main

import (
	"bufio"
	"fmt"
	"os"
)

type Monkey struct {
	name string

	isNumber bool
	number   int

	operator    byte
	leftMonkey  string
	rightMonkey string
}

func (m Monkey) Yell() int {
	if m.isNumber {
		return m.number
	}

	switch m.operator {
	case '+':
		return monkeys[m.leftMonkey].Yell() + monkeys[m.rightMonkey].Yell()
	case '-':
		return monkeys[m.leftMonkey].Yell() - monkeys[m.rightMonkey].Yell()
	case '*':
		return monkeys[m.leftMonkey].Yell() * monkeys[m.rightMonkey].Yell()
	case '/':
		return monkeys[m.leftMonkey].Yell() / monkeys[m.rightMonkey].Yell()
	}

	panic("Yell Not Reachable")
}

var (
	monkeys map[string]Monkey = make(map[string]Monkey)
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		name := line[:4]

		line = line[6:]
		if line[0] >= '0' && line[0] <= '9' {
			var number int
			fmt.Sscanf(line, "%d", &number)
			monkeys[name] = Monkey{name: name, isNumber: true, number: number}
		} else {
			var operator byte
			var leftMonkey, rightMonkey string
			fmt.Sscanf(line, "%s %c %s", &leftMonkey, &operator, &rightMonkey)
			monkeys[name] = Monkey{name: name, isNumber: false, leftMonkey: leftMonkey, operator: operator, rightMonkey: rightMonkey}
		}
	}

	// for i := range monkeys {
	// 	fmt.Println(monkeys[i])
	// }

	fmt.Println(monkeys["root"].Yell())
}
