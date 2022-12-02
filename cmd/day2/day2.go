package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input := readFile()

	fmt.Printf("Part1: %v \n", PartOne(input))
	fmt.Printf("Part2: %v \n", PartTwo(input))
}

// A for Rock, B for Paper, and C for Scissors
// X for Rock, Y for Paper, and Z for Scissors
// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)
func PartOne(input []string) int {
	totalScore := 0

	for _, round := range input {
		s := strings.Split(round, " ")
		elf := s[0]
		you := s[1]

		totalScore += getScore(you) + getOutcome(elf, you)
	}
	return totalScore
}

// A for Rock, B for Paper, and C for Scissors
// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)
func PartTwo(input []string) int {
	totalScore := 0
	for _, round := range input {
		s := strings.Split(round, " ")
		elf := s[0]
		outcome := s[1]

		you := decodeGesture(elf, outcome) // translate back to gestures used in part 1

		totalScore += getScore(you) + getOutcome(elf, you)
	}
	return totalScore
}

func readFile() []string {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(file), "\n")
}

func getScore(hand string) int {
	switch hand {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}

	log.Fatal("Invalid score input")
	return 0
}

func decodeGesture(elf string, you string) string { // to part 1 gestures
	if elf == "A" { // Rock
		if you == "X" { // you need to lose
			return "Z"
		}
		if you == "Y" { // you need to draw
			return "X"
		}
		if you == "Z" { // you need to win
			return "Y"
		}
	}
	if elf == "B" { // Paper
		if you == "X" { // you need to lose
			return "X"
		}
		if you == "Y" { // you need to draw
			return "Y"
		}
		if you == "Z" { // you need to win
			return "Z"
		}
	}
	if elf == "C" { // Seissors
		if you == "X" { // you need to lose
			return "Y"
		}
		if you == "Y" { // you need to draw
			return "Z"
		}
		if you == "Z" { // you need to win
			return "X"
		}
	}
	log.Fatal("Invalid gesture input")
	return "F"
}

func getOutcome(elf string, you string) int {
	if elf == "A" { // Rock
		if you == "X" {
			return 3
		}
		if you == "Y" {
			return 6
		}
		if you == "Z" {
			return 0
		}
	}

	if elf == "B" { // Paper
		if you == "X" {
			return 0
		}
		if you == "Y" {
			return 3
		}
		if you == "Z" {
			return 6
		}
	}

	if elf == "C" { // Scissors
		if you == "X" {
			return 6
		}
		if you == "Y" {
			return 0
		}
		if you == "Z" {
			return 3
		}
	}

	log.Fatal("Invalid gesture input")
	return 0
}
