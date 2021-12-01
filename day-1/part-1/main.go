package main

import (
	"advent-of-code-21/day-1/shared"
	"fmt"
)

func main() {
	increase := 0
	for i := 1; i < len(shared.Input); i++ {
		if shared.Input[i] > shared.Input[i-1] {
			increase++
		}
	}

	fmt.Println(increase)
}
