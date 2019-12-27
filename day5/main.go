package main

import (
	"github.com/harryganz/adventofcode/computer"
	"github.com/harryganz/adventofcode/utils"
	"os"
	"strconv"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	filepath := GOPATH + "/src/github.com/harryganz/adventofcode/data/day5.txt"

	inputStrings, err := utils.ScanFile(filepath, utils.SplitCommas)
	if err != nil {
		panic(err)
	}
	inputs := make([]int, 0)
	for _, v := range inputStrings {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		inputs = append(inputs, i)
	}

	_, err = computer.RunProgram(inputs)
	if err != nil {
		panic(err)
	}
}
