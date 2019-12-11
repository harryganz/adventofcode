package wires

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

// IsAt returns true if wire has a segment
// at x, y
func (w Wire) IsAt(x, y int) bool {
	return w.coordinates[x][y] == 1
}
