package utils

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ScanFileToInts parses input file and returns the contained data as a
// slice of ints
func ScanFileToInts(filepath string) ([]int, error) {
	result := make([]int, 0)
	var err error
	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	// Scan lines of file and parse integers for each line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return result, err
		}

		result = append(result, int(i))
	}

	return result, err
}

// ScanCommaSeparatedInts scans a file of comma separated ints to
// an int slice
func ScanCommaSeparatedInts(filepath string) ([]int, error) {
	result := make([]int, 0)
	var err error
	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	// Scan comma seperated values and parse integers for each value
	scanner := bufio.NewScanner(file)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// If at end of file and no data, return nothing
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		// Find index of comma, return what is in data, and advance past it
		if i := strings.Index(string(data), ","); i >= 0 {
			return i + 1, getNumeric(data[0:i]), nil
		}
		// If at end of file and there is more data, return the data that is left
		if atEOF {
			return len(data), getNumeric(data), nil
		}

		// Get more data
		return 0, nil, nil
	})
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return result, err
		}
		result = append(result, int(i))
	}
	return result, err
}

func getNumeric(data []byte) []byte {
	numericRegex := regexp.MustCompile("[0-9]+")
	return numericRegex.Find(data)
}
