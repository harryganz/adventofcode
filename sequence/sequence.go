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

func HasExactDouble(x int) bool {
	s := strconv.Itoa(x)

	nextToMap := make(map[byte]int)

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			if v, ok := nextToMap[s[i]]; ok {
				nextToMap[s[i]] = v + 1
			} else {
				nextToMap[s[i]] = 2
			}
		}
	}

	for _, v := range nextToMap {
		if v == 2 {
			return true
		}
	}

	return false
}
