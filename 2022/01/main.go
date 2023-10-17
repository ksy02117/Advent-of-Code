package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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
	var max [3]int
	sum := 0

	for {
		line, err := r.ReadString('\n')
		if err != nil && err == io.EOF {
			fmt.Println("end of file reached")
			break
		}
		line = strings.TrimSpace(line)
		// fmt.Println("[" + line + "]")

		if line == "" {
			if sum < max[2] {
				sum = 0
				continue
			}

			if sum < max[1] {
				max[2] = sum
				sum = 0
				continue
			}

			max[2] = max[1]
			if sum < max[0] {
				max[1] = sum
				sum = 0
				continue
			}

			max[1] = max[0]
			max[0] = sum
			sum = 0
		} else {
			str, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += str
		}
	}
	sum = 0
	for _, v := range max {
		sum += v
	}
	fmt.Println(max)
	fmt.Println(sum)
}
