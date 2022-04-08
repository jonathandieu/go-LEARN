package main

import (
	"fmt"
	"os"
)

func main() {
	reader := NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("You entered: ", input)
}
