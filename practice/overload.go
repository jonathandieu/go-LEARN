package main

import (
	"fmt"
)

// Takes in one or more ints as parameters
func Sum(nums ...int) int {
	sum := 0

	for _, number := range nums {
		sum += number
	}

	return sum
}

func main() {
	fmt.Println(Sum(6, 9))
}