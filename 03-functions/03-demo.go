//higher order functions

package main

import "fmt"

func main() {

	//function assigned to variable
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	//anonymous functions
	func() {
		fmt.Println("Printed from an anonymous function")
	}()

	/*
		func(x, y int) {
			fmt.Println("Subtract result from anonymous function", x-y)
		}(100, 200)
	*/
	subtractResult := func(x, y int) int {
		return x - y
	}(100, 200)

	fmt.Println("Subtract result from anonymous function", subtractResult)
}
