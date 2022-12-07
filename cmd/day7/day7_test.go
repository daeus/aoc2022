package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	verifyDeepEqual(t, 95437, PartOne(input))
}

func TestGetFileSystem(t *testing.T) {
	input := `$ cd /
dir a
1000 b.txt
2500 c.dat
$ cd a
$ ls
50 e.txt`

	actual := GetFileSystem(strings.Split(input, "\n"))

	file1 := File{
		Name: "b.txt",
		Size: 1000,
	}
	file2 := File{
		Name: "c.dat",
		Size: 2500,
	}
	file3 := File{
		Name: "e.txt",
		Size: 50,
	}

	folder1 := Node{
		Name: "a",
		Size: 50,
		Files: []File{
			file3,
		},
		Folders: []*Node{},
		Parent:  nil,
	}

	expected := Node{
		Name: "/",
		Size: 3550,
		Files: []File{
			file1,
			file2,
		},
		Folders: []*Node{&folder1},
	}

	verifyDeepEqual(t, expected.Name, actual.Name)
	verifyDeepEqual(t, expected.Size, actual.Size)
	verifyDeepEqual(t, len(expected.Files), len(actual.Files))
	verifyDeepEqual(t, len(expected.Folders), len(actual.Folders))
}

func verifyDeepEqual(t *testing.T, expected, actual interface{}) {
	res := reflect.DeepEqual(expected, actual)
	if !res {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
