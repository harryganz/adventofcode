package fuel

import "testing"

func TestCalculateFuel(t *testing.T) {
	table := []struct {
		in     int
		expect int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, c := range table {
		got := CalculateFuel(c.in)
		if got != c.expect {
			t.Errorf("for CalculateMass(%d) got %d, expected %d", c.in, got, c.expect)
		}
	}
}

func TestCalculateTotalFuel(t *testing.T) {
	table := []struct {
		in     int
		expect int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
		{2, 0},
	}

	for _, c := range table {
		got := CalculateTotalFuel(c.in)
		if got != c.expect {
			t.Errorf("CalculateTotalFuel(%d) got %d, expected %d", c.in, got, c.expect)
		}
	}
}
