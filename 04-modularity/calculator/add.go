package calculator

import "fmt"

func init() {
	fmt.Println("calculator is initialized")
}

func Add(x, y int) int {
	increment()
	return x + y
}
