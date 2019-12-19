package computer

import "errors"

func RunProgram(input []int) ([]int, error) {
	for pos := 0; pos < len(input) && input[pos] != 99; pos += 4 {
		instruction := parseInstruction(input[pos])

		out, err := execute(input, instruction, input[pos+1:pos+4])
		if err != nil {
			return input, err
		}
		input = out
	}

	return input, nil
}

func execute(input, instruction, params []int) ([]int, error) {
	paramValues := make([]int, len(params))
	for i, param := range params {
		switch mode := instruction[i+1]; mode {
		case 0:
			paramValues[i] = input[param]
		case 1:
			paramValues[i] = param
		}
	}
	switch opCode := instruction[0]; opCode {
	case 1:
		loc := paramValues[2]
		x := paramValues[0]
		y := paramValues[1]
		input[loc] = x + y
	case 2:
		loc := paramValues[2]
		x := paramValues[0]
		y := paramValues[1]
		input[loc] = x * y
	default:
		return []int{}, errors.New("Unknown opcode")
	}

	return input, nil
}

func parseInstruction(instruction int) []int {
	mode1 := instruction / 10000
	mode2 := (instruction - mode1*10000) / 1000
	mode3 := (instruction - mode1*10000 - mode2*1000) / 100
	opCode := instruction - mode1*10000 - mode2*1000 - mode3*100

	return []int{opCode, mode1, mode2, mode3}
}
