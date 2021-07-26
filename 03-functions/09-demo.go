package main

import (
	"errors"
	"fmt"
)

func main() {
	q, r, err := divide(100, 0)
	if err != nil {
		fmt.Println("Something went wrong!", err)
		return
	}
	fmt.Println(q, r)
}

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("Invalid argument(s). Cannot divide by zero")
		return
	}
	quotient, remainder = x/y, x%y
	return
}
