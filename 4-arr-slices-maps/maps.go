package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7 // add a kv pair
	m["k2"] = 13
	delete(m, "k2") // delete a kv pair

	val, ok := m["k2"] // check if a key exists
	fmt.Println("1. val:", val, "| ok:", ok)

	// declare and init
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("2. map: ", n)

}
