package main

import (
	"fmt"
)

// Takes in one or more ints as parameters
// This is a variadic function that can take any number of ints as arguments
func Sum(nums ...int) int {
	sum := 0

	for _, number := range nums {
		sum += number
	}

	return sum
}

// func ListyInterface(nums ...interface{}) {

// 	for _, number := range nums {
// 		sum += number
// 	}
// }

func main() {
	fmt.Println(Sum(6, 9))
}
