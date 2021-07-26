/*
	create increment() and decrement() functions which will increment and decrement the value of the same variable.
	The outcome of increment and decrement should not be able to be influenced by any other code
*/

/* increment() // => 1
increment() // => 2
increment() // => 3
increment() // => 4
decrement() //=> 3
decrement() //=> 2
decrement() //=> 1
decrement() //=> 0 */

package main

import "fmt"

func main() {
	increment, decrement := getCounters()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
}

func getCounters() (func() int, func() int) {
	count := 0
	increment := func() int {
		count++
		return count
	}
	decrement := func() int {
		count--
		return count
	}
	return increment, decrement
}
