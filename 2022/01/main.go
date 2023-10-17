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
	var max [3]int
	ch := make(chan int, 1)

	go getCal(ch)

	for v := range ch {
		if v >= max[0] {
			max[2], max[1], max[0] = max[1], max[0], v
			continue
		}

		if v >= max[1] {
			max[2], max[1] = max[1], v
			continue
		}

		if v >= max[2] {
			max[2] = v
			continue
		}
	}

	sum := 0
	for _, v := range max {
		sum += v
	}
	fmt.Println(max)
	fmt.Println(sum)
}

func getCal(ch chan int) {
	file, _ := os.Open("input.txt")
	defer file.Close()
	r := bufio.NewReader(file)
	sum := 0

	for {
		line, err := r.ReadString('\n')
		if err != nil && err == io.EOF {
			fmt.Println("end of file reached")
			close(ch)
			return
		}
		line = strings.TrimSpace(line)

		if line == "" {
			fmt.Println(sum)
			ch <- sum
			sum = 0
		} else {
			str, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += str
		}
	}
}

/*
func sol1() {
	file, _ := os.Open("input.txt")
	defer file.Close()

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
*/
