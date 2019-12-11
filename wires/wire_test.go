package wires

import (
	"reflect"
	"testing"
)

func TestAddSegment(t *testing.T) {
	w1 := New()
	w1.AddSegment("R", 2)
	w1.AddSegment("D", 1)
	w1.AddSegment("L", 3)
	w1.AddSegment("U", 1)

	table := []struct {
		x      int
		y      int
		expect bool
	}{
		{0, 0, true},
		{1, 0, true},
		{2, 0, true},
		{3, 0, false},
		{2, -1, true},
		{2, -2, false},
		{1, -1, true},
		{0, -1, true},
		{-1, -1, true},
		{-2, -1, false},
		{-1, 0, true},
		{-1, 1, false},
	}

	for _, c := range table {
		if got := w1.IsAt(c.x, c.y); got != c.expect {
			t.Errorf("IsAt(%d, %d) returned %t. Expected %t\n", c.x, c.y, got, c.expect)
		}
	}
}

func TestGetIntersections(t *testing.T) {
	w1 := New()
	w1.AddSegment("R", 5)
	w1.AddSegment("U", 4)

	w2 := New()
	w2.AddSegment("U", 2)
	w2.AddSegment("R", 6)
	w2.AddSegment("U", 1)
	w2.AddSegment("L", 1)

	got := w1.GetIntersections(w2)
	expect := map[int]map[int]int{
		0: {0: 1},
		5: {2: 1, 3: 1},
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Expected: %v, Got: %v\n", expect, got)
	}
}

func TestparseSegment(t *testing.T) {
	table := []struct {
		in string
		d  string
		l  int
	}{
		{"R5", "R", 5},
		{"U100", "U", 100},
		{"", "R", 0},
		{"U", "R", 0},
	}

	for _, c := range table {
		d, l := parseSegment(c.in)

		if c.d != d || c.l != l {
			t.Errorf("parseSegment(%s) got: %s, %d. expected: %s, %d\n", c.in, d, l, c.d, c.l)
		}
	}
}

func TestClosestIntersectionDistance(t *testing.T) {
	table := []struct {
		w1     []string
		w2     []string
		expect int
	}{
		{
			[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			159,
		},
	}

	for _, c := range table {
		w1 := New()
		w2 := New()
		w1.AddSegments(c.w1)
		w2.AddSegments(c.w2)
		got := ClosestIntersectionDistance(w1.GetIntersections(w2))
		if c.expect != got {
			t.Errorf("Did not get expected closest intersection. got: %d, expected: %d\n", got, c.expect)
		}
	}
}
