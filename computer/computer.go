package computer

import (
	"errors"
)

func RunProgram(input []int) ([]int, error) {
	output := make([]int, len(input))
	copy(output, input)
	var err error
	var register int
	var advance int
	for pos := 0; pos < len(output) && output[pos] != 99; pos += advance {
		instruction := parseInstruction(output[pos])
		register, advance, err = execute(output, instruction, pos, register)
		if err != nil {
			return output, err
		}
	}

	return output, nil
}

func execute(input, instruction []int, pos, register int) (int, int, error) {
	paramModes := instruction[1:]
	getParam := paramGetter(pos, input, paramModes)

	switch opcode := instruction[0]; opcode {
	case 1:
		x1, x2 := getParam(0), getParam(1)
		input[input[pos+3]] = x1 + x2
		return register, 4, nil
	case 2:
		x1, x2 := getParam(0), getParam(1)
		input[input[pos+3]] = x1 * x2
		return register, 4, nil
	case 3:
		value := getParam(0)
		return value, 2, nil
	case 4:
		input[input[pos+1]] = register
		return register, 2, nil
	default:
		return 0, 0, errors.New("Unknown opcode")
	}
}

func parseInstruction(instruction int) []int {
	mode1 := instruction / 10000
	mode2 := (instruction - mode1*10000) / 1000
	mode3 := (instruction - mode1*10000 - mode2*1000) / 100
	opCode := instruction - mode1*10000 - mode2*1000 - mode3*100

	return []int{opCode, mode1, mode2, mode3}
}

func paramGetter(pos int, input []int, mode []int) func(int) int {
	return func(index int) int {
		if mode[index] == 0 {
			return input[input[pos+index+1]]
		} else {
			return input[pos+index+1]
		}
	}
}
