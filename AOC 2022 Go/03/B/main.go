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
		// sort them so your can use two way comparison for dup checking
		item1 := translate(sc.Text())
		slices.Sort(item1)

		sc.Scan()
		item2 := translate(sc.Text())
		slices.Sort(item2)

		sc.Scan()
		item3 := translate(sc.Text())
		slices.Sort(item3)

		// only advance smallest array
		for i, j, k := 0, 0, 0; i < len(item1) && j < len(item2) && k < len(item3); {
			if item1[i] > item2[j] {
				j++
				continue
			}
			if item1[i] < item2[j] {
				i++
				continue
			}
			if item1[i] > item3[k] {
				k++
				continue
			}
			if item1[i] < item3[k] {
				i++
				j++
				continue
			}
			fmt.Println(item1[i], item2[j], item3[k])
			sum += int(item1[i])
			break
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
