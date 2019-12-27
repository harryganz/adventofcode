package computer

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
)

var defaultReader *bufio.Reader = bufio.NewReader(os.Stdin)
var defaultWriter *bufio.Writer = bufio.NewWriter(os.Stdout)

func RunProgram(input []int) ([]int, error) {
	output := make([]int, len(input))
	copy(output, input)
	var err error
	var advance int
	for pos := 0; pos < len(output) && output[pos] != 99; pos += advance {
		instruction := parseInstruction(output[pos])
		advance, err = execute(output, instruction, pos)
		if err != nil {
			return output, err
		}
	}

	return output, nil
}

func execute(input, instruction []int, pos int) (int, error) {
	paramModes := instruction[1:]
	getParam := paramGetter(pos, input, paramModes)

	switch opcode := instruction[0]; opcode {
	case 1:
		x1, x2 := getParam(0), getParam(1)
		input[input[pos+3]] = x1 + x2
		return 4, nil
	case 2:
		x1, x2 := getParam(0), getParam(1)
		input[input[pos+3]] = x1 * x2
		return 4, nil
	case 3:
		s, err := defaultReader.ReadString('\n')
		if err != nil {
			return 2, err
		}
		value, err := strconv.Atoi(trimEndline(s))
		if err != nil {
			return 2, err
		}
		input[input[pos+1]] = value
		return 2, nil
	case 4:
		value := getParam(0)
		_, err := defaultWriter.WriteString(strconv.Itoa(value))
		if err != nil {
			return 2, err
		}
		return 2, nil
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

func paramGetter(pos int, input []int, mode []int) func(int) int {
	return func(index int) int {
		if mode[index] == 0 {
			return input[input[pos+index+1]]
		} else {
			return input[pos+index+1]
		}
	}
}

func trimEndline(s string) string {
	rg := regexp.MustCompile("\r?\n")
	return rg.ReplaceAllString(s, "")
}
