package utils

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// ScansFile scans the file given using bufio.Scanner and
// the splitFunc provided, returning a slice of strings of each
// token
func ScanFile(filepath string, splitFunc bufio.SplitFunc) ([]string, error) {
	result := make([]string, 0)
	var err error
	file, err := os.Open(filepath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		s := scanner.Text()
		result = append(result, s)
	}
	return result, err
}

func cleanNewLine(data []byte) []byte {
	nonNewLineRegex := regexp.MustCompile("[^\n\r\v]+")
	return nonNewLineRegex.Find(data)
}

// Splits function that splits by commans and returns tokens with
// whitepace stripped
func SplitCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// If at end of file and no data, return nothing
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	// Find index of comma, return what is in data, and advance past it
	if i := strings.Index(string(data), ","); i >= 0 {
		return i + 1, cleanNewLine(data[0:i]), nil
	}
	// If at end of file and there is more data, return the data that is left
	if atEOF {
		return len(data), cleanNewLine(data), nil
	}

	// Get more data
	return 0, nil, nil
}
