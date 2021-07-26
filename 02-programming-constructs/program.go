package main

import "fmt"

func main() {
	/* Programming constructs -> if, switch, for, func */

	/* if construct */
	/*
		no := 21
		if no%2 == 0 {
			fmt.Println(no, "is an even number")
		} else {
			fmt.Println(no, "is an odd number")
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Println(no, "is an even number")
	} else {
		fmt.Println(no, "is an odd number")
	}

	//for construct
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//imitating the 'while' using 'for'
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println(numSum)

	//indefinite loop
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Println(x)

}
