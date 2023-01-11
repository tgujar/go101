package main

func plus(a, b int) int {
	return a + b
}

func makeMultiplier(mul int) (int, func(s int) int) {
	defer (func() { println("makeMultiplier() is done") })()

	anon_func := func() {
		println("I'm an anonymous function")
	}
	defer anon_func()

	multiplier := func(a int) int {
		return a * mul
	}

	return mul, multiplier
}

func main() {
	plus(1, 2)
	mul, x2 := makeMultiplier(2) // defer executes LIFO
	println("mul:", mul)
	println("x2(3):", x2(3))
}
