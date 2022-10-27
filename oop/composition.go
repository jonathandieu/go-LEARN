package main

import (
	"fmt"
)

type alive struct {

}

type walkable struct {

}

type swimmable struct {

}

type duck struct {
	a alive
	w walkable
	s swimmable
}

func (alive) eat() {
	fmt.Println("Alive and eating")
}
func (alive) sleep() {
	fmt.Println("Alive and sleeping")
}
func (alive) walk() {
	fmt.Println("Alive and strolling through")
}

func (d duck) eat() {
	d.a.eat()
}
func (d duck) sleep() {
	d.a.sleep()
}
func (d duck) walk() {
	d.a.walk()
}
func main() {
	fmt.Println("Hello, world!")
	d := duck{}
	d.eat()
	d.sleep()
	d.walk()
}