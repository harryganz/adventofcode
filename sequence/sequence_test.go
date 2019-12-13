package sequence

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	f := func(start, end int) []int {
		return Filter(start, end, []func(int) bool{HasDouble, HasIncreasingDigits})
	}

	table := []struct {
		start  int
		end    int
		expect []int
	}{
		{10, 20, []int{11}},
		{20, 40, []int{22, 33}},
		{1, 10, []int{}},
	}

	for _, c := range table {
		if got := f(c.start, c.end); !reflect.DeepEqual(c.expect, got) {
			t.Errorf("for %d - %d, expected: %v, got: %v", c.start, c.end, c.expect, got)
		}
	}
}

func TestHasDouble(t *testing.T) {
	table := []struct {
		in     int
		expect bool
	}{
		{11, true},
		{123, false},
		{12234, true},
		{12344, true},
	}

	for _, c := range table {
		if got := HasDouble(c.in); got != c.expect {
			t.Errorf("HasDouble(%d) returned %t, expected %t\n", c.in, got, c.expect)
		}
	}
}

func TestHasIncreasingDigits(t *testing.T) {
	table := []struct {
		in     int
		expect bool
	}{
		{123, true},
		{10, false},
		{1123, true},
		{1254, false},
		{111, true},
	}

	for _, c := range table {
		if got := HasIncreasingDigits(c.in); got != c.expect {
			t.Errorf("HasIncreasingDigits(%d) returned %t, expected %t\n", c.in, got, c.expect)
		}
	}
}
