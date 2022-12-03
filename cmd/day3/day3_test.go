package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	PmmdzqPrVvPwwTWBwg
	wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	ttgJtRGJQctTZtZT
	CrZsJsPPZsGzwwsLwLmpwMDw`

	actual := PartOne(input)
	expect := 157
	verifyEqual(t, expect, actual)
}

func TestPartTwo(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
	jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	PmmdzqPrVvPwwTWBwg
	wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	ttgJtRGJQctTZtZT
	CrZsJsPPZsGzwwsLwLmpwMDw`

	actual := PartTwo(input)
	expect := 70
	verifyEqual(t, expect, actual)
}

func verifyEqual(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
