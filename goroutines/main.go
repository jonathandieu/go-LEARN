package main

import (
	"fmt"
	"time"
)

//
func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println("Time: ", i)
	}
}

func main() {
	var someString string = "Welcome to Jon's introduction to concurrency in Go!"
	var args string

	fmt.Println(someString)
	go compute(5)

	go compute(5)
	fmt.Println("Second time's the charm...")

	fmt.Println("Please type in a string")
	fmt.Scan(&args)
	fmt.Println("Ayooo! You typed:", args)
}
