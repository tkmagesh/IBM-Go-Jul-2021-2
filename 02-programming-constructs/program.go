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

	//switch construct
	/*
		newNo := 20
		switch newNo % 2 {
		case 0:
			fmt.Println(newNo, "is an even number")
		case 1:
			fmt.Println(newNo, "is an odd number")
		}
	*/

	newNo := 20
	switch {
	case newNo%2 == 0:
		fmt.Println(newNo, "is an even number")
	case newNo%2 == 1:
		fmt.Println(newNo, "is an odd number")
	}

	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Good"
		score 10  => "Excellent"
		otherwise => "Unknown score"
	*/

	score := 7
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Unknown score")
	}

	//fallthrough construct
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
		fallthrough
	case 7:
		fmt.Println("n <= 7")
	case 8:
		fmt.Println("n <= 8")
		fallthrough
	case 9:
		fmt.Println("n <= 9")
		fallthrough
	case 10:
		fmt.Println("n <= 10")
	}
}
