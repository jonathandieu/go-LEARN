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
	fmt.Println(someString)
	go compute(5)

	go compute(5)
	fmt.Println("Second time's the charm...")
	fmt.Scanln()
}
