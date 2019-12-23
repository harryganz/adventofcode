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
		register, advance, err = execute(output, instruction, output[pos+1:pos+4], register)
		if err != nil {
			return output, err
		}
	}

	return output, nil
}

func execute(input, instruction, params []int, register int) (int, int, error) {
	paramValues := make([]int, 3)
	// Populate paramValues based on parameter mode
	for i := 0; i < len(params); i++ {
		switch mode := instruction[i+1]; mode {
		case 0:
			paramValues[i] = input[params[i]]
		case 1:
			paramValues[i] = params[i]
		}
	}

	switch opcode := instruction[0]; opcode {
	case 1:
		x1, x2 := paramValues[0], paramValues[1]
		value := x1 + x2
		input[params[2]] = value
		return value, 4, nil
	case 2:
		x1, x2 := paramValues[0], paramValues[1]
		value := x1 * x2
		input[params[2]] = value
		return value, 4, nil
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
