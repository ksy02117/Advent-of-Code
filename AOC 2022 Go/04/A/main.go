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

	cnt := 0
	for sc.Scan() {
		line := sc.Text()

		var s1, s2, e1, e2 int
		fmt.Sscanf(line, "%d-%d,%d-%d", &s1, &e1, &s2, &e2)

		if s1 >= s2 && e1 <= e2 {
			cnt++
			continue
		}

		if s1 <= s2 && e1 >= e2 {
			cnt++
			continue
		}
	}

	fmt.Println(cnt)
}
