package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	items []int

	operator byte
	operand  string

	test    int
	trueTo  int
	falseTo int

	cnt int
}

func (m *Monkey) Operate() {
	for i := range m.items {
		// Operate
		m.cnt++
		operand := 0

		switch m.operand {
		case "old":
			operand = m.items[i]
		default:
			operand, _ = strconv.Atoi(m.operand)
		}

		switch m.operator {
		case '+':
			m.items[i] = (m.items[i] + operand) / 3
		case '*':
			m.items[i] = m.items[i] * operand / 3
		}

		// Test and Throw
		if m.items[i]%m.test == 0 {
			monkeys[m.trueTo].items = append(monkeys[m.trueTo].items, m.items[i])
		} else {
			monkeys[m.falseTo].items = append(monkeys[m.falseTo].items, m.items[i])
		}
	}

	m.items = make([]int, 0)
}

var (
	monkeys []Monkey
)

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		if !strings.HasPrefix(line, "Monkey") {
			continue
		}
		var monkey Monkey

		// Get Starting Items
		sc.Scan()
		line = sc.Text()
		line = line[18:]
		itemsString := strings.Split(line, ", ")
		monkey.items = make([]int, 0)
		for i := range itemsString {
			v, _ := strconv.Atoi(itemsString[i])
			monkey.items = append(monkey.items, v)
		}

		// Get Operation
		sc.Scan()
		line = sc.Text()
		fmt.Sscanf(line, "  Operation: new = old %c %s", &monkey.operator, &monkey.operand)

		// Get Test
		sc.Scan()
		line = sc.Text()
		fmt.Sscanf(line, "  Test: divisible by %d", &monkey.test)

		// Get True Monkey
		sc.Scan()
		line = sc.Text()
		fmt.Sscanf(line, "    If true: throw to monkey %d", &monkey.trueTo)

		// Get False Monkey
		sc.Scan()
		line = sc.Text()
		fmt.Sscanf(line, "    If false: throw to monkey %d", &monkey.falseTo)

		sc.Scan()
		monkeys = append(monkeys, monkey)
	}

	for k := 0; k < 20; k++ {
		for i := range monkeys {
			monkeys[i].Operate()
		}
	}
	// for i := range items {
	// 	for j := range items[i] {
	// 		// monkeys[i].cnt++
	// 		monkey, monkeyAfter := i, 0
	// 		value := items[i][j]
	// 		for k := 0; k < 20; k++ {
	// 			monkeys[monkey].cnt++
	// 			monkeyAfter, value = monkeys[monkey].Operate(value)
	// 			if monkeyAfter > monkey {
	// 				k--
	// 			}
	// 			monkey = monkeyAfter
	// 		}
	// 	}
	// }

	for i := range monkeys {
		fmt.Println(monkeys[i])
	}

	max1, max2 := 0, 0
	for i := range monkeys {
		if max1 < monkeys[i].cnt {
			max2, max1 = max1, monkeys[i].cnt
		} else if max2 < monkeys[i].cnt {
			max2 = monkeys[i].cnt
		}
	}

	fmt.Println(max1 * max2)
}
