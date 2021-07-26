package main

import "fmt"

func main() {
	x := 10
	y := 20
	wrappedAdd := wrapLog(add)
	wrappedSub := wrapLog(subtract)
	fmt.Println("Addition:", wrappedAdd(x, y))
	fmt.Println("Subtraction:", wrappedSub(x, y))
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func wrapLog(fn func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		fmt.Println("Before invocation")
		result := fn(x, y)
		fmt.Println("After invocation")
		return result
	}
}
