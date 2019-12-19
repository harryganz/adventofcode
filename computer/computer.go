package computer

import "errors"

func RunProgram(input []int) ([]int, error) {
	for pos := 0; pos < len(input) && input[pos] != 99; pos += 4 {
		instruction := parseInstruction(input[pos])
		x1 := input[input[pos+1]]
		x2 := input[input[pos+2]]
		loc := input[pos+3]
		out, err := execute(instruction[0], x1, x2)
		if err != nil {
			return input, err
		}
		input[loc] = out
	}

	return input, nil
}

func execute(opcode, x1, x2 int) (int, error) {
	switch opcode {
	case 1:
		return x1 + x2, nil
	case 2:
		return x1 * x2, nil
	default:
		return 0, errors.New("Unknown opcode")
	}
}

func parseInstruction(instruction int) []int {
	mode1 := instruction / 10000
	mode2 := (instruction - mode1*10000) / 1000
	mode3 := (instruction - mode1*10000 - mode2*1000) / 100
	opCode := instruction - mode1*10000 - mode2*1000 - mode3*100

	return []int{opCode, mode1, mode2, mode3}
}
