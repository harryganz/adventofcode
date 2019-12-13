package main

import (
	"bufio"
	"fmt"
	"github.com/harryganz/adventofcode/utils"
	"github.com/harryganz/adventofcode/wires"
	"os"
	"strings"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	filepath := GOPATH + "/src/github.com/harryganz/adventofcode/data/day3.txt"

	inputLines, err := utils.ScanFile(filepath, bufio.ScanLines)
	if err != nil {
		panic(err)
	}

	line1 := inputLines[0]
	line2 := inputLines[1]

	w1 := wires.New()
	w2 := wires.New()

	w1.AddSegments(strings.Split(line1, ","))
	w2.AddSegments(strings.Split(line2, ","))

	intersections := w1.GetIntersections(w2)

	distance := wires.ClosestIntersectionManhattan(intersections)

	fmt.Printf("Closest intersection to port = %d\n", distance)

	distance2 := wires.ClosestIntersectionLinear(intersections)

	fmt.Printf("Closest intersections along wires = %d\n", distance2)
}
