package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Stack struct {
	items []byte
}

func main() {
	input := readFile()
	partOneAns := PartOne(input)
	partTwoAns := PartTwo(input)
	fmt.Printf("Part1: %v \n", partOneAns)
	fmt.Printf("Part2: %v \n", partTwoAns)
}

func PartOne(input string) string {
	crates := GetCrate(input)

	commands := strings.Split(input, "\n")
	for _, c := range commands {
		if len(c) == 0 || string(c[0]) != "m" {
			continue
		}

		numOfItem, from, to := decodeCommand(c)

		items := crates[from-1].pop(numOfItem)
		for i := 0; i < numOfItem; i++ {
			crates[to-1].append(items[numOfItem-i-1])
		}
	}
	return getTopStack(crates)
}

func PartTwo(input string) string {
	crates := GetCrate(input)
	commands := strings.Split(input, "\n")

	for _, c := range commands {
		if len(c) == 0 || string(c[0]) != "m" {
			continue
		}

		numOfItem, from, to := decodeCommand(c)

		items := crates[from-1].pop(numOfItem)
		crates[to-1].appendAll(items)
	}

	return getTopStack(crates)
}

func decodeCommand(c string) (int, int, int) {
	var label string
	var numOfItem, from, to int
	fmt.Sscan(c, &label, &numOfItem, &label, &from, &label, &to)

	return numOfItem, from, to
}

func getTopStack(s []Stack) string {
	res := ""
	for _, crate := range s {
		if len(crate.items) > 0 {
			res += string(crate.items[len(crate.items)-1])
		}
	}
	return res
}

//---- Stack functions ----
func (s *Stack) toStr() []string {
	str := []string{}
	for _, item := range s.items {
		str = append(str, string(item))
	}
	return str
}

func (s *Stack) unshift(item byte) {
	s.items = append([]byte{item}, s.items...)
}

func (s *Stack) pop(length int) []byte {
	last := s.items[len(s.items)-length:]
	s.items = s.items[:len(s.items)-length]
	return last
}

func (s *Stack) appendAll(item []byte) {
	s.items = append(s.items, item...)
}

func (s *Stack) append(item byte) {
	s.items = append(s.items, item)
}

func GetCrate(input string) []Stack {
	crates := []Stack{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if string(line[1]) == "1" {
			break
		}

		for i := 1; i < len(line); i += 4 {
			// dynamically increase the size of crates
			if i/4 >= len(crates) {
				crates = append(crates, Stack{})
			}

			if string(line[i]) != " " {
				crates[i/4].unshift(line[i])
			}
		}
	}
	return crates
}

func readFile() string {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
