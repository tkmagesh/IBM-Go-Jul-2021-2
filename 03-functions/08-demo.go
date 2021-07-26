package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		fmt.Println("deffered from main function")
		if r := recover(); r != nil {
			fmt.Println("Something went wrong [in main]!")
			fmt.Println(r)
		}
	}()
	fmt.Println("Invoking the divide function")
	fmt.Println(divideClient(100, 0))
	fmt.Println("End of main")
}

func divideClient(x, y int) (quotient, remainder int) {
	defer func() {
		fmt.Println("deffered from divideClient function")
		if r := recover(); r != nil {
			fmt.Println("Recovering from divide function")
			fmt.Println("Something went wrong!")
			fmt.Println(r)
			debug.PrintStack()
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

func divide(x, y int) (int, int) {
	defer fmt.Println("deffered from divide function")
	if y == 0 {
		panic("Invalid arguments. Cannot divide by zero")
	}
	return x / y, x % y
}
