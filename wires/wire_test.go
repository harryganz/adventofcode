package wires

import "testing"

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
