package main

import "fmt"

func main() {
	//addResult := add(100, 200)
	var addResult = add(100, 200)
	fmt.Println(addResult)

	//fmt.Println(divide(100, 7))
	/*
		quotient, remainder := divide(100, 7)
		fmt.Println(quotient, remainder)
	*/
	quotient, _ := divide(100, 7)
	fmt.Println(quotient)
}

/*
func add(x int, y int) int {
	return x + y
}
*/

/*
func add(x, y int) int {
	return x + y
}
*/

func add(x, y int) (result int) {
	result = x + y
	return
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
