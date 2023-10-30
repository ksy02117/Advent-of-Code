package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ksy02117/Advent-of-Code/2022/07/tree"
)

// builds out the folder structure and calculate everything after building folder structure
// since i am not sure if the input revisits folders multiple times
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

	// 70000000 - root.size + x = 30000000
	answer := getAnswer(&root, 30000000-70000000+root.Size)
	fmt.Println(answer.Name, answer.Size)

}

func getAnswer(n *tree.Folder, size int) *tree.Folder {
	var out *tree.Folder = nil

	if n.Size >= size {
		out = n
	}

	for i := range n.Folders {
		v := getAnswer(&n.Folders[i], size)
		if v == nil {
			continue
		}

		if out == nil || out.Size > v.Size {
			out = v
		}
	}

	return out
}
