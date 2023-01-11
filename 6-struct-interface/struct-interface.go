package main

import "fmt"

// ref: https://gobyexample.com/interfaces

// We dont have classes, we instead only have structs and interfaces

// Interfaces are a collection of method signatures
type geometry interface {
	area() float64
	perim() float64
}

// Structs are a collection of typed fields
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// interfaces are composable
type object interface {
	geometry
	color() string
}

// structs are composable
type box struct {
	rect
	color string
}

func main() {

	r := rect{width: 3, height: 4}
	fmt.Println("1. Measure r:")
	measure(r) // we can pass in r because struct rect implements the interface geometry

	b := box{rect{width: 3, height: 4}, "red"}
	fmt.Println("\n2. Box height: ", b.height)

}
