package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	finput := readFile()
	input := strings.Split(finput, "\n")

	elves := []int{}
	temp := 0
	for _, c := range input {
		if c == "" {
			elves = append(elves, temp)
			temp = 0
		} else {
			temp += toInt(string(c))
		}
	}

	sort.SliceStable(elves, func(i, j int) bool {
		return elves[j] < elves[i]
	})

	fmt.Printf("Part1: %v \n", elves[0])
	fmt.Printf("Part2: %v \n", elves[0]+elves[1]+elves[2])
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func readFile() string {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}
