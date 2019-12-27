package computer

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestRunProgram(t *testing.T) {
	defaultReader = bufio.NewReader(strings.NewReader("1\r\n"))
	table := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{3, 0, 4, 0, 99}, []int{1, 0, 4, 0, 99}},
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
