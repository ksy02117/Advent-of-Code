package main

import (
	"bufio"
	"fmt"
	"os"
)

type Snafu struct {
	digits []int
}

func (n Snafu) String() string {
	out := ""
	for i := len(n.digits) - 1; i >= 0; i-- {
		out += string(toSnafu[n.digits[i]])
	}
	return out
}

func ToSnafu(s string) Snafu {
	out := Snafu{digits: make([]int, 0)}
	for i := len(s) - 1; i >= 0; i-- {
		out.digits = append(out.digits, toInt[s[i]])
	}
	return out
}

func (n *Snafu) Add(n2 Snafu) {
	roundUp := 0
	out := Snafu{make([]int, 0)}
	for i := 0; i < len(n.digits) || i < len(n2.digits) || roundUp != 0; i++ {
		digitSum := roundUp
		if i < len(n.digits) {
			digitSum += n.digits[i]
		}
		if i < len(n2.digits) {
			digitSum += n2.digits[i]
		}

		if digitSum > 2 {
			digitSum -= 5
			roundUp = 1
		} else if digitSum < -2 {
			digitSum += 5
			roundUp = -1
		} else {
			roundUp = 0
		}

		if digitSum < -2 || digitSum > 2 {
			panic("")
		}
		out.digits = append(out.digits, digitSum)
	}

	n.digits = out.digits
}

var (
	toInt   map[byte]int = map[byte]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}
	toSnafu map[int]byte = map[int]byte{-2: '=', -1: '-', 0: '0', 1: '1', 2: '2'}
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	number := Snafu{make([]int, 0)}
	for sc.Scan() {
		number.Add(ToSnafu(sc.Text()))
	}

	fmt.Println(number)
}
