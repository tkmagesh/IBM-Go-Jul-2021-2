package main

import "fmt"

func main() {
	/* defer func() {
		fmt.Println("Exiting from main")
	}() */
	defer fmt.Println("Exiting from main")
	fmt.Println("Main started")
	f1()
	fmt.Println("Main completed")
}

func f1() {
	defer fmt.Println("[defer - df1] exiting from f1")
	defer fmt.Println("[defer - df2] exiting from f1")
	defer fmt.Println("[defer - df3] exiting from f1")
	fmt.Println("f1 invoked")
	msg := f2()
	fmt.Println("f1 received message:", msg)
}

func f2() (message string) {
	defer func() {
		fmt.Println("[defer - df1] exiting from f2")
		message = "message from defer-df1 of f2 func"
	}()
	defer fmt.Println("[defer - df2] exiting from f2")
	defer fmt.Println("[defer - df3] exiting from f2")
	fmt.Println("f2 invoked")
	message = "message from the f2 func"
	return
}
