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

	// if left is true number then bool is true and leftNumber gets a value
	// otherwise, leftMonkey gets a value
	isLeftNumber bool
	leftMonkey   string
	leftNumber   int

	operator byte

	// same logic as left number
	isRightNumber bool
	rightMonkey   string
	rightNumber   int
}

func (m Monkey) String() string {
	out := m.name
	out += ": "
	if m.isNumber {
		out += fmt.Sprint(m.number)
	} else {
		if m.isLeftNumber {
			out += fmt.Sprint(m.leftNumber)
		} else {
			out += m.leftMonkey
		}
		out += fmt.Sprintf(" %c ", m.operator)
		if m.isRightNumber {
			out += fmt.Sprint(m.rightNumber)
		} else {
			out += m.rightMonkey
		}
	}

	return out
}

func (m *Monkey) Yell(monkeys *map[string]*Monkey) (int, bool) {
	if m.isNumber {
		return m.number, m.name != "humn"
	}

	// only calculate when each side is not a true value
	if !m.isLeftNumber {
		m.leftNumber, m.isLeftNumber = (*monkeys)[m.leftMonkey].Yell(monkeys)
	}
	if !m.isRightNumber {
		m.rightNumber, m.isRightNumber = (*monkeys)[m.rightMonkey].Yell(monkeys)
	}

	switch m.operator {
	case '+':
		m.number = m.leftNumber + m.rightNumber
	case '-':
		m.number = m.leftNumber - m.rightNumber
	case '*':
		m.number = m.leftNumber * m.rightNumber
	case '/':
		m.number = m.leftNumber / m.rightNumber
	case '=':
		if m.leftNumber == m.rightNumber {
			m.number = 1
		} else {
			m.number = 0
		}
	}

	// if the both ends of the numbers are true value (not affected by humn) then turn the monkey into number monkey
	if m.isLeftNumber && m.isRightNumber {
		m.isNumber = true
		return m.number, true
	}

	return m.number, false
}

// after getting all the true values, convert the monkeys into calculating what human is supposed to say
func (m Monkey) Convert() (string, *Monkey) {
	if m.name == "root" {
		if m.isLeftNumber {
			return m.rightMonkey, &Monkey{name: m.rightMonkey, isNumber: true, number: m.leftNumber}
		} else {
			return m.leftMonkey, &Monkey{name: m.leftMonkey, isNumber: true, number: m.rightNumber}
		}
	}

	if m.operator == '/' {
		// abc : def / 123          -> def : abc * 123
		return m.leftMonkey, &Monkey{name: m.leftMonkey, leftMonkey: m.name, operator: '*', isRightNumber: true, rightNumber: m.rightNumber}
	}

	if m.operator == '-' {
		if m.isLeftNumber {
			// abc : 123 - def      -> def : 123 - abc
			return m.rightMonkey, &Monkey{name: m.rightMonkey, isLeftNumber: true, leftNumber: m.leftNumber, operator: '-', rightMonkey: m.name}
		} else {
			// abc : def - 123      -> def : abc + 123
			return m.leftMonkey, &Monkey{name: m.leftMonkey, leftMonkey: m.name, operator: '+', isRightNumber: true, rightNumber: m.rightNumber}
		}
	}

	if m.operator == '+' {
		if m.isLeftNumber {
			// abc : 123 + def      -> def : abc - 123
			return m.rightMonkey, &Monkey{name: m.rightMonkey, leftMonkey: m.name, operator: '-', isRightNumber: true, rightNumber: m.leftNumber}
		} else {
			// abc : def + 123      -> def : abc - 123
			return m.leftMonkey, &Monkey{name: m.leftMonkey, leftMonkey: m.name, operator: '-', isRightNumber: true, rightNumber: m.rightNumber}
		}
	}

	if m.operator == '*' {
		if m.isLeftNumber {
			// abc : 123 * def      -> def : abc / 123
			return m.rightMonkey, &Monkey{name: m.rightMonkey, leftMonkey: m.name, operator: '/', isRightNumber: true, rightNumber: m.leftNumber}
		} else {
			// abc : def + 123      -> def : abc - 123
			return m.leftMonkey, &Monkey{name: m.leftMonkey, leftMonkey: m.name, operator: '/', isRightNumber: true, rightNumber: m.rightNumber}
		}
	}

	return "", &Monkey{}
}

var (
	monkeys  map[string]*Monkey = make(map[string]*Monkey)
	monkeys2 map[string]*Monkey = make(map[string]*Monkey)
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
			monkeys[name] = &Monkey{name: name, isNumber: true, number: number}
		} else {
			var operator byte
			var leftMonkey, rightMonkey string
			fmt.Sscanf(line, "%s %c %s", &leftMonkey, &operator, &rightMonkey)
			if name == "root" {
				monkeys[name] = &Monkey{name: name, isNumber: false, leftMonkey: leftMonkey, operator: '=', rightMonkey: rightMonkey}
			} else {
				monkeys[name] = &Monkey{name: name, isNumber: false, leftMonkey: leftMonkey, operator: operator, rightMonkey: rightMonkey}
			}
		}
	}

	// for i := range monkeys {
	// 	fmt.Println(monkeys[i])
	// }

	monkeys["root"].Yell(&monkeys)

	// for i := range monkeys {
	// 	if !monkeys[i].isNumber {
	// 		fmt.Println(*monkeys[i])
	// 		name, monkey := monkeys[i].Convert()
	// 		monkeys2[name] = monkey
	// 	}
	// }

	fmt.Println()

	// for i := range monkeys2 {
	// 	fmt.Println(monkeys2[i])
	// }

	fmt.Println(monkeys2["humn"].Yell(&monkeys2))

	// for i := range monkeys2 {
	// 	fmt.Println(*monkeys2[i])
	// }
}
