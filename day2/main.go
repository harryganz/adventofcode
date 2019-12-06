package main

import (
	"fmt"
	"github.com/harryganz/adventofcode/computer"
	"github.com/harryganz/adventofcode/utils"
	"os"
	"strconv"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	filepath := GOPATH + "/src/github.com/harryganz/adventofcode/data/day2.txt"

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

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			currentInput := make([]int, len(inputs))
			copy(currentInput, inputs)
			currentInput[1] = noun
			currentInput[2] = verb
			outputs, err := computer.RunProgram(currentInput)
			if outputs[0] == 19690720 && err == nil {
				fmt.Printf("noun: %d, verb: %d, 100 * noun + verb = %d\n", noun, verb, 100*noun+verb)
				break
			}
		}
	}
}
