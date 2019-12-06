package main

import (
	"bufio"
	"fmt"
	"github.com/harryganz/adventofcode/fuel"
	"github.com/harryganz/adventofcode/utils"
	"os"
	"strconv"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	filepath := GOPATH + "/src/github.com/harryganz/adventofcode/data/day1.txt"
	inputStrings, err := utils.ScanFile(filepath, bufio.ScanLines)
	if err != nil {
		panic(err)
	}
	input := make([]int, 0)
	for _, v := range inputStrings {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		input = append(input, i)
	}

	sum := 0
	for _, v := range input {
		sum += fuel.CalculateTotalFuel(v)
	}
	fmt.Printf("Fuel required for dry mass is %d\n", sum)
}
