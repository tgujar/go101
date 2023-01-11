package main

import "fmt"

func main() {

	for n := 0; n <= 5; n++ { // for loop c style
		// if-else chain
		if n%2 == 0 {
			fmt.Println("even")
		} else if n == 3 {
			fmt.Println("three")
		} else {
			fmt.Println("odd")
		}
	}

	// while loop
	n := 0
	for n < 5 {
		fmt.Println(n)
		n++
	}

}
