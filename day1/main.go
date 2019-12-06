package main

import (
	"fmt"
	"github.com/harryganz/adventofcode/fuel"
	"github.com/harryganz/adventofcode/utils"
	"os"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	filepath := GOPATH + "/src/github.com/harryganz/adventofcode/data/day1.txt"
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
