package fuel

import "math"

// CalculateFuel returns the amount of fuel required
// for a given mass
func CalculateFuel(mass int) int {
	return int(math.Floor(float64(mass)/3.0) - 2.0)
}

// Calculates the total fuel needed for a dry mass
// including the other fuel needed
func CalculateTotalFuel(mass int) int {
	sum := 0
	for fc := CalculateFuel(mass); fc > 0; fc = CalculateFuel(fc) {
		sum += fc
	}
	return sum
}
