package main

import (
	"fmt"
	"github.com/harryganz/adventofcode/fuel"
	"github.com/harryganz/adventofcode/utils"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("missing input file path")
	}
	filepath := args[1]
	input, err := utils.ScanFileToInts(filepath)
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, v := range input {
		sum += fuel.CalculateTotalFuel(v)
	}
	fmt.Printf("Fuel required for dry mass is %d\n", sum)
}
