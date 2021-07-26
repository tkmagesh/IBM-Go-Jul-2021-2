package main

import "fmt"

func main() {
	x := 10
	y := 20
	compute(x, y, add)
	compute(x, y, subtract)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func compute(x, y int, fn func(int, int) int) {
	result := fn(x, y)
	fmt.Printf("Result of processing %d and %d is %d\n", x, y, result)
}
