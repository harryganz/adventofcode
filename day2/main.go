package main

import (
	"fmt"
	"github.com/harryganz/adventofcode/computer"
	"github.com/harryganz/adventofcode/utils"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("missing filepath")
	}
	filepath := args[1]

	inputs, err := utils.ScanCommaSeparatedInts(filepath)
	if err != nil {
		panic(err)
	}

	output := 0
	noun := 0
	verb := 0
	for ; noun < 100 && output != 19690720; noun++ {
		for ; verb < 100 && output != 19690720; verb++ {
			currentInputs := inputs
			currentInputs[1] = noun
			currentInputs[2] = verb
			outputs, err := computer.RunProgram(currentInputs)
			if err == nil {
				output = outputs[0]
				fmt.Println(output)
			}
		}
	}
	noun -= 1
	verb -= 1
	fmt.Printf("noun: %d, verb: %d, 100 * noun + verb = %d", noun, verb, 100*noun+verb)
}
