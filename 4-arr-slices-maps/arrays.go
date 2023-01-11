package main

import "fmt"

func main() {
	var a [5]int // arrays are fixed length
	fmt.Println("Empty arrays are initialized to zero values:", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Declare and init:", b)
}
