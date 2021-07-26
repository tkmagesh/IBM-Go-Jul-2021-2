//variadic functions
package main

import "fmt"

func main() {
	fmt.Println(sum(10, 20, 30))
	fmt.Println(sum(10, 20, 30, 40))
	fmt.Println(sum(10, 20, 30, 40, 50))
}

func sum(nos ...int) int {
	//nos => slice of int (collection)
	result := 0
	/* for idx := 0; idx < len(nos); idx++ {
		result += nos[idx]
	} */

	for _, no := range nos {
		result += no
	}
	return result
}
