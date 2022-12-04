package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := readFile()
	partOneAns := PartOne(input)
	PartTwoAns := PartTwo(input)

	log.Printf("Part 1: %v", partOneAns)
	log.Printf("Part 2: %v", PartTwoAns)
}

func PartOne(input string) int {
	elvesPairs := strings.Split(input, "\n")
	pairs := 0
	for _, p := range elvesPairs {
		elves := strings.Split(p, ",")

		elf_A := stringsToInts(strings.Split(elves[0], "-"))
		elf_B := stringsToInts(strings.Split(elves[1], "-"))

		if isRangeSubset(elf_A, elf_B) || isRangeSubset(elf_B, elf_A) {
			pairs++
		}
	}
	return pairs
}

func PartTwo(input string) int {
	elvesPairs := strings.Split(input, "\n")
	pairs := 0
	for _, p := range elvesPairs {
		elves := strings.Split(p, ",")

		elf_A := stringsToInts(strings.Split(elves[0], "-"))
		elf_B := stringsToInts(strings.Split(elves[1], "-"))

		if isOverlap(elf_A, elf_B) {
			pairs++
		}
	}
	return pairs
}

func isOverlap(setA []int, setB []int) bool {
	if isInRange(setA, setB[0]) || isInRange(setA, setB[1]) || isInRange(setB, setA[0]) || isInRange(setB, setA[1]) {
		return true
	}
	return false
}

func isInRange(r []int, dig int) bool {
	if r[1] >= dig && r[0] <= dig {
		return true
	}
	return false
}

func isRangeSubset(mainSet []int, subSet []int) bool {
	if mainSet[1] >= subSet[1] && subSet[0] >= mainSet[0] {
		return true
	}
	return false
}

func stringsToInts(input []string) []int {
	output := []int{}
	for _, i := range input {
		output = append(output, toInt(i))
	}
	return output
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
