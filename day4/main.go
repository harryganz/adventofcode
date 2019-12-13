package main

import (
	"fmt"
	"github.com/harryganz/adventofcode/sequence"
)

func main() {
	valid := sequence.Filter(273025, 767253, []func(int) bool{sequence.HasDouble, sequence.HasIncreasingDigits})

	fmt.Println("Valid numbers in sequence = ", len(valid))
}
