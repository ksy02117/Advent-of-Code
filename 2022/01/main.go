package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(file)
	max, sum := 0, 0

	for {
		line, err := r.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		line = line[:len(line)-1]

		if line == "" {
			fmt.Println("new elf")
			if sum >= max {
				max = sum
			}
			sum = 0
		} else {
			str, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%v\n", str)
			sum += str
		}
	}

	fmt.Println(max)
}
