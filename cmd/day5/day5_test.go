package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetCrate(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3`
	actual := cratesToString(GetCrate(input))
	fmt.Println(actual)
	expected := map[int][]string{}
	expected[0] = []string{"Z", "N"}
	expected[1] = []string{"M", "C", "D"}
	expected[2] = []string{"P"}

	verifyDeepEqual(t, expected, actual)
}

func TestStack(t *testing.T) {
	s := Stack{}
	s2 := Stack{}
	s2.append([]byte("H")[0])
	s.append([]byte("A")[0])
	s.append([]byte("B")[0])
	s.append([]byte("C")[0])
	s.append([]byte("D")[0])

	verifyDeepEqual(t, []string{"A", "B", "C", "D"}, s.toStr())

	items := s.pop(2)

	verifyDeepEqual(t, []string{"A", "B"}, s.toStr())

	s2.appendAll(items)

	verifyDeepEqual(t, []string{"H", "C", "D"}, s2.toStr())

	s.unshift([]byte("Z")[0])

	verifyDeepEqual(t, []string{"Z", "A", "B"}, s.toStr())
}

func TestDecodeCommand(t *testing.T) {
	input := "move 1 from 2 to 3"
	num, from, to := decodeCommand(input)

	verifyDeepEqual(t, 1, num)
	verifyDeepEqual(t, 2, from)
	verifyDeepEqual(t, 3, to)
}

func cratesToString(crates []Stack) map[int][]string {
	str := map[int][]string{}
	for i, crate := range crates {
		str[i] = append(str[i], crate.toStr()...)
	}
	return str
}

func verifyDeepEqual(t *testing.T, expected, actual interface{}) {
	res := reflect.DeepEqual(expected, actual)
	if !res {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
