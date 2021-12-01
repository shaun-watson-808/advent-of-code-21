package main

import (
	"advent-of-code-21/day-1/shared"
	"fmt"
)

func main() {
	increase := 0
	previousSum := 0
	for i := 0; i < len(shared.Input)-2; i++ {
		sum := shared.Input[i] + shared.Input[i+1] + shared.Input[i+2]

		if i != 0 {
			if sum > previousSum {
				increase++
			}
		}

		previousSum = sum
	}

	fmt.Println(increase)
}
