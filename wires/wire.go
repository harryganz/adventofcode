package wires

import (
	"strconv"
)

// Wire stores the coordinates of a wire and the
// ending point of its last segment
type Wire struct {
	coordinates map[int]map[int]int
	endX        int
	endY        int
}

// New creates a new wire object
func New() *Wire {
	coords := make(map[int]map[int]int)
	endX := 0
	endY := 0

	coords[0] = make(map[int]int)
	coords[0][0] = 1
	return &Wire{coords, endX, endY}
}

func parseSegment(segment string) (string, int) {
	d, l := "R", 0
	defer func() {
		if r := recover(); r != nil {
			d, l = "R", 0
		}
	}()

	d = string(segment[0])
	l, err := strconv.Atoi(segment[1:])
	if err != nil {
		panic(err)
	}
	return d, l
}

// AddSegment adds a segment to the existing wire
func (w *Wire) AddSegment(direction string, length int) {
	var isHorizontal bool
	var sign int

	switch direction {
	case "R":
		isHorizontal = true
		sign = 1
	case "L":
		isHorizontal = true
		sign = -1
	case "U":
		isHorizontal = false
		sign = 1
	case "D":
		isHorizontal = false
		sign = -1
	default:
		panic("unknown direction")
	}

	for i := 0; i < length; i++ {
		if isHorizontal {
			w.endX += sign
		} else {
			w.endY += sign
		}
		if _, ok := w.coordinates[w.endX]; !ok {
			w.coordinates[w.endX] = make(map[int]int)
		}
		w.coordinates[w.endX][w.endY] = 1
	}
}

// AddSegments adds multiple segments from an array of strings
func (w *Wire) AddSegments(s []string) {
	for _, segment := range s {
		d, l := parseSegment(segment)
		w.AddSegment(d, l)
	}
}

// IsAt returns true if wire has a segment
// at x, y
func (w Wire) IsAt(x, y int) bool {
	return w.coordinates[x][y] == 1
}

// GetIntersections returns a map of all the points where
// this wire intersects with another wire
func (w Wire) GetIntersections(o *Wire) map[int]map[int]int {
	result := make(map[int]map[int]int)
	for x, col := range w.coordinates {
		for y := range col {
			if _, ok := o.coordinates[x][y]; ok {
				if _, ok := result[x]; !ok {
					result[x] = make(map[int]int)
				}
				result[x][y] = 1
			}
		}
	}

	return result
}

// ClosestIntersectionDistance returns the closest intersection's distance
// given a map of intersections
func ClosestIntersectionDistance(intersections map[int]map[int]int) int {
	minDistance := 0
	for x, col := range intersections {
		for y := range col {
			if x != 0 && y != 0 {
				if minDistance == 0 {
					minDistance = x + y
				} else if minDistance > (x + y) {
					minDistance = x + y
				}
			}
		}
	}

	return minDistance
}
