package computer

import "errors"

func RunProgram(input []int) ([]int, error) {
	for pos := 0; pos < len(input) && input[pos] != 99; pos += 4 {
		instruction := parseInstruction(input[pos])
		err := execute(input, instruction, input[pos+1:pos+4])
		if err != nil {
			return input, err
		}
	}

	return input, nil
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
