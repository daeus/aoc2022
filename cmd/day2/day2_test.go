package main

import "testing"

func TestPartOne(t *testing.T) {
	case1 := []string{
		"A Y",
	}
	case2 := []string{
		"B X",
	}
	case3 := []string{
		"C Z",
	}

	actual1 := PartOne(case1)
	actual2 := PartOne(case2)
	actual3 := PartOne(case3)

	verifyEqual(t, 8, actual1)
	verifyEqual(t, 1, actual2)
	verifyEqual(t, 6, actual3)
}

func TestPartTwo(t *testing.T) {
	case1 := []string{
		"A Y",
	}
	case2 := []string{
		"B X",
	}
	case3 := []string{
		"C Z",
	}

	actual1 := PartTwo(case1)
	actual2 := PartTwo(case2)
	actual3 := PartTwo(case3)

	verifyEqual(t, 4, actual1)
	verifyEqual(t, 1, actual2)
	verifyEqual(t, 7, actual3)
}

func verifyEqual(t *testing.T, expected, actual int) {
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
