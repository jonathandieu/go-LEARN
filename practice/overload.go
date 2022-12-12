package main

import (
	"fmt"
	"errors"
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

func PersonInterface(args ...interface{}) {
	// mandatory parameters
	var name string // left uninitialized

	// optional parameters
	var ssn int = -1
	var phoneNumber int = -1
	var favoriteColor string = ""

	// Make sure we have at least the min number of parameters
	if len(args) < 1 {
		panic("Not enough parameters provided!")
	}

	for i, p := range args {
		switch i {
		case 0:
			param, ok := p.(string)
			if !ok {
				panic("First parameter provided not of type string!")
			}
			name = param

		case 1:
			param, ok := p.(int)
			if !ok {
				panic("First parameter provided not of type int!")
			}
			ssn = param
		case 2:
			param, ok := p.(int)
			if !ok {
				panic("First parameter provided not of type int!")
			}
			phoneNumber = param
		case 3:
			param, ok := p.(string)
			if !ok {
				panic("First parameter provided not of type string!")
			}
			favoriteColor = param
		}
	}
}

func main() {
	fmt.Println(Sum(6, 9))
}

func Player(args ...interface{}) {

    name, x, y, color, err := playerParams(args...)
    if nil != err {
        panic(err.Error())
    }

    // ...
}

func playerParams(args ...interface{}) (name string, x int, y int, color string, err error) {

    // We initialize each of the optional parameters to their default value.
    x = 0            // ← We initialize x to zero (0) here.
    y = 0            // ← We initialize y to zero (0) here.
    color = "green"  // ← We initialize color to "green" here.

    // Since we have 1 mandatory parameter, make sure the number of parameters
    // we have is at least 1.
    //
    // We can figure out the number of parameters passed to us with len(args).
    if 1 > len(args) {
        err = errors.New("Not enough parameters.")
        return
    }


    // Get any parameters passed to us out of the args variable into "real"
    // variables we created for them.
    for i,p := range args {
        switch i {
            case 0: // name
                param, ok := p.(string)
                if !ok {
                    err = errors.New("1st parameter not type string.")
                    return
                }
                name = param

            case 1: // x
                param, ok := p.(int)
                if !ok {
                    err = errors.New("2nd parameter not type int.")
                    return
                }
                x = param

            case 2: // y
                param, ok := p.(int)
                if !ok {
                    err = errors.New("2nd parameter not type int.")
                    return
                }
                y = param

            case 3: // color
                param, ok := p.(string)
                if !ok {
                    err = errors.New("3rd parameter not type string.")
                    return
                }
                color = param

            default:
                err = errors.New("Wrong parameter count.")
                return
        }
    }


    return
}