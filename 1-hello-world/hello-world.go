// Entry point is main() function in package main

/* A complete program is created by linking a single, unimported package called the main package
 with all the packages it imports, transitively. The main package must have package name main and
declare a function main that takes no arguments and returns no value. */

package main

import "fmt" // read more at https://pkg.go.dev/fmt

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("Welcome to %s%d \n", "CSE", 224)
}

// Compile and create executable -> go build {filename.go}
// Compile and run (binary not saved) -> go run {filename.go}
