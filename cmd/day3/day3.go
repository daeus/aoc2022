package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input := readFile()

	score1 := PartOne(input)
	score2 := PartTwo(input)
	log.Printf("Part1: %v \n", score1)
	log.Printf("Part2: %v \n", score2)

}

func PartOne(input string) int {
	rucksacks := strings.Split(input, "\n")
	score := 0
	for _, rucksack := range rucksacks {
		size := len(rucksack)
		c1 := rucksack[0 : size/2]
		c2 := rucksack[size/2 : size]

		hashC1 := make(map[string]bool)
		intersections := []string{}

		for _, item := range c1 {
			hashC1[string(item)] = true
		}

		for _, item := range c2 {
			if hashC1[string(item)] {
				intersections = append(intersections, string(item))
			}
		}
		score += getPriorityScore(intersections[0])
	}

	return score
}

func PartTwo(input string) int {
	rucksacks := strings.Split(input, "\n")

	score := 0

	for i := 0; i < len(rucksacks); i += 3 {

		hashElf_1 := make(map[string]bool)
		hashElf_2 := make(map[string]bool)
		intersections := []string{}

		for _, item := range rucksacks[i] {
			hashElf_1[string(item)] = true
		}

		for _, item := range rucksacks[i+1] {
			hashElf_2[string(item)] = true
		}

		for _, item := range rucksacks[i+2] {
			if hashElf_1[string(item)] && hashElf_2[string(item)] {
				intersections = append(intersections, string(item))
			}
		}
		fmt.Printf("Intersections: %v \n", intersections)
		score += getPriorityScore(intersections[0])
	}
	return score
}

func getPriorityScore(item string) int {
	// get alphabet ASCII
	// subtract 64 to get A=1, B=2, C=3, etc
	if (item >= "A") && (item <= "Z") {
		return int(item[0]) - 38

		// assume lowercase
	} else {
		return int(item[0]) - 96
	}
}

func readFile() string {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return string(file)
}
