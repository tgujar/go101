package main

import (
	"fmt"
	"math"
)

func declareAndInitialize() {
	// Declaring variables
	var k int // declare (zero-valued by default)

	var i int = 10            // declare and initialize (explicit)
	const s = "Hello, World!" // declare and initialize (type inference)
	b := true                 // declare and initialize (short)
	fmt.Println(i, k, s, b)   // need this because unused "vars" are an error for Golang
}

func untypedConstants() {
	// Untyped constants
	const a = 5
	b := 5
	fmt.Printf("With const: %v\n", a*math.Pi)
	fmt.Printf("Without const: %v\n", float64(b)*math.Pi) // type conversion has to be explicit

	var d int8 = 25
	d = d * 8

	fmt.Println(d) // int overflow

	// untyped const can do arbitrary precision math
}

func main() {
	declareAndInitialize()
	untypedConstants()
}

// More types -> https://go.dev/ref/spec#Types
