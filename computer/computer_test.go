package computer

import (
	"reflect"
	"testing"
)

func TestRunProgram(t *testing.T) {
	table := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{[]int{3, 5, 4, 0, 99, -1}, []int{-1, 5, 4, 0, 99, -1}},
		{[]int{10003, 5, 4, 0, 99}, []int{5, 5, 4, 0, 99}},
	}

	for _, c := range table {
		got, err := RunProgram(c.input)
		if err != nil {
			t.Fatalf("Error running test %s\n", err)
		}
		if !reflect.DeepEqual(got, c.output) {
			t.Errorf("RunProgram(%d) returned %d, expected %d\n", c.input, got, c.output)
		}
	}
}

func TestParseInstruction(t *testing.T) {
	table := []struct {
		in     int
		expect []int
	}{
		{2, []int{2, 0, 0, 0}},
		{11, []int{11, 0, 0, 0}},
		{101, []int{1, 0, 0, 1}},
		{1102, []int{2, 0, 1, 1}},
		{10003, []int{3, 1, 0, 0}},
	}

	for _, c := range table {
		if got := parseInstruction(c.in); !reflect.DeepEqual(got, c.expect) {
			t.Errorf("parstInstruction(%d). got: %v, expected: %v\n", c.in, got, c.expect)
		}
	}
}
