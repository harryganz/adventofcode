package computer

import (
	"errors"
)

func RunProgram(input []int) ([]int, error) {
	output := make([]int, len(input))
	copy(output, input)
	for pos := 0; pos < len(output) && output[pos] != 99; pos += 4 {
		instruction := parseInstruction(output[pos])
		err := execute(output, instruction, output[pos+1:pos+4])
		if err != nil {
			return output, err
		}
	}

	return output, nil
}

func execute(input, instruction, params []int) error {
	switch opcode := instruction[0]; opcode {
	case 1:
		x1, x2 := input[params[0]], input[params[1]]
		input[params[2]] = x1 + x2
	case 2:
		x1, x2 := input[params[0]], input[params[1]]
		input[params[2]] = x1 * x2
	default:
		return errors.New("Unknown opcode")
	}

	return nil
}

func parseInstruction(instruction int) []int {
	mode1 := instruction / 10000
	mode2 := (instruction - mode1*10000) / 1000
	mode3 := (instruction - mode1*10000 - mode2*1000) / 100
	opCode := instruction - mode1*10000 - mode2*1000 - mode3*100

	return []int{opCode, mode1, mode2, mode3}
}
