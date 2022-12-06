package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input := readFile()
	fmt.Printf("Part1: %v \n", PartOne(input))
	fmt.Printf("Part2: %v \n", PartTwo(input))
}

func PartOne(input string) int {
	return getMarker(input, 4)
}

func PartTwo(input string) int {
	return getMarker(input, 14)
}

func getMarker(input string, size int) int {
	hashed := map[string]int{}
	from := 0

	for i, c := range input {
		if _, ok := hashed[string(c)]; ok {
			// check if the char is in the current window
			if hashed[string(c)] >= from {
				// move to next non repeating char
				from = hashed[string(c)] + 1
			}
		}
		// update index to the latest
		hashed[string(c)] = i

		if i-from+1 == size {
			return i + 1
		}
	}
	return 0
}

func readFile() string {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(file)
}
