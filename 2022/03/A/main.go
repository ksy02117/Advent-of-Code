package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	sum := 0
	for sc.Scan() {
		line := sc.Text()

		item1 := translate(line[:len(line)/2])
		item2 := translate(line[len(line)/2:])
		slices.Sort(item1)
		slices.Sort(item2)

		for i, j := 0, 0; i < len(item1) && j < len(item2); {
			if item1[i] > item2[j] {
				j++
			} else if item1[i] < item2[j] {
				i++
			} else {
				fmt.Println(item1[i], item2[j])
				sum += int(item1[i])
				break
			}
		}
	}
	fmt.Println(sum)
}

func translate(input string) []byte {
	output := []byte(input)

	for i, v := range output {
		if v >= 'a' && v <= 'z' {
			output[i] -= 'a' - 1
		} else {
			output[i] -= 'A' - 27
		}
	}

	return output
}
