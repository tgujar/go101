package main

import "fmt"

func main() {
	// slices are dynamic length

	// make([]type, len, cap) allocates a slice of length len and capacity cap with type "type"
	s := make([]string, 2) // also zero valued by default
	s[0] = "a"
	s[1] = "b"
	s = append(s, "c", "d", "e", "f") // append() is a built-in function
	fmt.Println("1. len:", len(s), "| cap:", cap(s), "| s:", s)

	// Slice operation
	t := s[2:5]
	t[0] = "z"
	fmt.Println("2. len:", len(t), "| cap:", cap(t), "| t:", t)

	fmt.Println("3. s after t is modifies:", s) // What can be said from this?

	//	Capacity is the number of elements in the underlying array,
	//	counting from the first element in the slice.

	// Slicing beyond len
	t = t[:cap(t)]
	fmt.Println("4. Extending beyond slice", "len:", len(t), "| cap:", cap(t), "| t:", t)

	// Copy a slice
	c := make([]string, len(s))
	copy(c, s)

	// 2d slices
	//twoD := make([][]int, 3)
}
