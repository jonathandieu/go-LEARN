package main

import (
	"fmt"
)

type alive struct {
}

func (alive) eat() {
	fmt.Println("Alive and eating")
}

type walkable struct {
}

func (alive) walk() {
	fmt.Println("Alive and strolling through")
}

type swimmable struct {
}

func (alive) swim() {
	fmt.Println("Alive and swimming")
}

type flyable struct {
}

type duck struct {
	a alive
	w walkable
	s swimmable
}

func (d duck) eat() {
	d.a.eat()
}
func (d duck) sleep() {
	d.a.swim()
}
func (d duck) walk() {
	d.a.walk()
}

type goose struct {
	alive
	walkable
	swimmable
	flyable
}

func (goose) walk() {
	fmt.Println("Goose walkin.")
}

func main() {
	fmt.Println("Hello, world!")
	d := duck{}
	d.eat()
	d.sleep()
	d.walk()

	g := goose{}
	g.walk()
}
