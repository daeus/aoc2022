package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Node struct {
	Name    string
	Files   []File
	Folders []*Node
	Parent  *Node
	Size    int
}

type File struct {
	Name string
	Size int
}

func (n *Node) appendFile(name string, size int) {
	f := File{Name: name, Size: size}
	n.Files = append(n.Files, f)
	updateParentSize(n, size)
}

func (n *Node) appendFolder(name string) {
	f := Node{Name: name, Parent: n, Size: 0, Files: []File{}, Folders: []*Node{}}
	n.Folders = append(n.Folders, &f)
}

func (n *Node) findFolder(name string) (*Node, bool) {
	for _, f := range n.Folders {
		if f.Name == name {
			return f, true
		}
	}
	return nil, false
}

func main() {
	input := readFile()

	fmt.Printf("Part 1: %v\n", PartOne(input))
	fmt.Printf("Part 2: %v\n", PartTwo(input))
}

func PartOne(input string) int {
	commands := strings.Split(input, "\n")
	sys := GetFileSystem(commands)
	return sizeOfSmallFolder(&sys)
}

func PartTwo(input string) int {
	commands := strings.Split(input, "\n")
	sys := GetFileSystem(commands)

	totalDiskSize := 70000000
	unusedSizeNeeded := 30000000

	currentUnusedSize := totalDiskSize - sys.Size

	neededSize := unusedSizeNeeded - currentUnusedSize

	fmt.Printf(" . Total Disk Size: %v\n", totalDiskSize)
	fmt.Printf(" . Current Unused Size: %v\n", currentUnusedSize)
	fmt.Printf(" . Needed Size: %v\n", neededSize)
	fmt.Printf(" . root size: %v\n", sys.Size)

	return findSmallestPossibleSize(&sys, neededSize, sys.Size)

}

func findSmallestPossibleSize(node *Node, neededSize int, smallestPossible int) int {
	if node.Size >= neededSize && node.Size < smallestPossible {
		smallestPossible = node.Size
	}

	for _, f := range node.Folders {
		smallestPossible = findSmallestPossibleSize(f, neededSize, smallestPossible)
	}

	return smallestPossible
}

func sizeOfSmallFolder(node *Node) int {
	limit := 100000
	currentSize := 0

	if limit >= node.Size {
		currentSize += node.Size
	}

	for _, f := range node.Folders {
		currentSize += sizeOfSmallFolder(f)
	}

	return currentSize
}

func GetFileSystem(cs []string) Node {
	root := Node{
		Name:    "/",
		Files:   []File{},
		Folders: []*Node{},
		Parent:  nil,
		Size:    0,
	}

	currentNode := &root

	for _, c := range cs {
		if isCommand(c) {
			if ok, loc := isCd(c); ok {
				if loc == ".." {
					currentNode = currentNode.Parent
				} else {
					if node, ok := currentNode.findFolder(loc); ok {
						currentNode = node
					}
				}
			}

			if isLs(c) {
				// do nothing
			}
		} else {
			if ok, name := isFolder(c); ok {
				currentNode.appendFolder(name)
			} else {
				size, name := getFileDetails(c)
				currentNode.appendFile(name, size)
			}
		}
	}
	return root
}

func updateParentSize(node *Node, size int) {
	node.Size += size
	if node.Parent != nil {
		updateParentSize(node.Parent, size)
	}
}

func isFolder(c string) (bool, string) {
	if strings.HasPrefix(c, "dir") {
		return true, strings.Split(c, " ")[1]
	}
	return false, ""
}

func getFileDetails(c string) (int, string) {
	s := strings.Split(c, " ")
	size, _ := strconv.Atoi(s[0])
	return size, s[1]
}

func isCommand(c string) bool {
	return strings.HasPrefix(c, "$")
}

func isCd(c string) (bool, string) {
	if strings.HasPrefix(c, "$ cd") {
		return true, strings.Split(c, " ")[2]
	}
	return false, ""
}

func isLs(c string) bool {
	return strings.HasPrefix(c, "$ ls")
}

func readFile() string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
