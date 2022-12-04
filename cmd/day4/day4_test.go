package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `2-4,6-8
	2-3,4-5
	5-7,7-9
	2-8,3-7
	6-6,4-6
	2-6,4-8`

	actual := PartOne(input)
	expect := 2
	verifyEqual(t, expect, actual)
}

func TestPartTwo(t *testing.T) {
}

func verifyEqual(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
