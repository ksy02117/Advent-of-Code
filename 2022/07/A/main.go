package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ksy02117/Advent-of-Code/2022/07/tree"
)

func main() {
	file, _ := os.Open("../input.txt")
	defer file.Close()
	sc := bufio.NewScanner(file)

	root := tree.Folder{
		Name:    "/",
		Folders: make([]tree.Folder, 0),
		Files:   make([]tree.File, 0),
		Parent:  nil,
		Size:    0,
	}
	currentDir := &root
	var name string

	for sc.Scan() {
		line := sc.Text()

		if line == "$ cd /" {
			currentDir = &root
			continue
		}

		if line == "$ cd .." {
			currentDir = currentDir.MoveToParent()
			continue
		}

		if strings.HasPrefix(line, "$ cd") {
			fmt.Sscanf(line, "$ cd %s", &name)
			currentDir = currentDir.MoveTo(name)
			// fmt.Print("current Directory: ")
			// fmt.Println(currentDir)
			continue
		}

		if strings.HasPrefix(line, "$ ls") {
			continue
		}

		if strings.HasPrefix(line, "dir") {
			fmt.Sscanf(line, "dir %s", &name)
			currentDir.AddFolder(name)
			// fmt.Printf("%s : %s\n", currentDir.Name, name)
			continue
		}

		var size int
		fmt.Sscanf(line, "%d %s", &size, &name)
		currentDir.AddFile(name, size)
	}

	root.CalculateSize()

	answer := getAnswer(&root)
	fmt.Println(answer)
}

func getAnswer(n *tree.Folder) int {
	sum := 0
	for _, folder := range n.Folders {
		sum += getAnswer(&folder)
	}

	if n.Size <= 100000 {
		return sum + n.Size
	}

	return sum
}
