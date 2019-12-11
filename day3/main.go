package main

import (
	"fmt"
	"github.com/harryganz/adventofcode/utils"
	"github.com/harryganz/adventofcode/wires"
	"os"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	filepath := GOPATH + "/src/github.com/harryganz/adventofcode/data/day3.txt"

	inputStrings, err := utils.ScanFile(filepath, utils.SplitCommas)
	if err != nil {
		panic(err)
	}

	fmt.Println(inputStrings[0:3])

	w1 := wires.New()
	w1.AddSegment("R", 2)
	w1.AddSegment("D", 1)
	w1.AddSegment("L", 3)
	w1.AddSegment("U", 1)
}
