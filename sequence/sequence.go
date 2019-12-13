package sequence

import "strconv"

func Filter(start int, end int, cbs []func(int) bool) []int {
	result := make([]int, 0)
	for i := start; i <= end; i++ {
		inResult := true
		for _, cb := range cbs {
			inResult = inResult && cb(i)
		}
		if inResult {
			result = append(result, i)
		}
	}

	return result
}

func HasDouble(x int) bool {
	s := strconv.Itoa(x)

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}

func HasIncreasingDigits(x int) bool {
	s := strconv.Itoa(x)

	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}

	return true
}
