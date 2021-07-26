package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("deffered from main function")
		if r := recover(); r != nil {
			fmt.Println("Something went wrong [in main]!")
			fmt.Println(r)
		}
	}()
	fmt.Println("Invoking the divideClient function")
	q, r, err := divideClient(100, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(q, r)
	fmt.Println("End of main")
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		fmt.Println("deffered from divideClient function")
		if r := recover(); r != nil {
			fmt.Println("Recovering from divide function")
			fmt.Println("Something went wrong!")
			fmt.Println(r)
			//debug.PrintStack()
			err = errors.New("Invalid argument(s)!")
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
