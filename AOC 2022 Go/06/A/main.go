package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	sc.Scan()
	line := sc.Text()

	fmt.Println(len(line))

	s, e := 0, 0

	for e = range line {
		for i := s; i < e; i++ {
			// if repeat happens we know the packet does not form until after the first value
			if line[i] == line[e] {
				fmt.Printf("%v(%c) : %v(%c)\n", i, line[i], e, line[e])
				s = i + 1
				break
			}
		}

		if e-s+1 == 4 {
			break
		}
	}

	fmt.Println(e + 1)
}
