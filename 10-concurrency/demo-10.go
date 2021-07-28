package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go print("Hello", 3, ch1, ch2)
	go print("World", 2, ch2, ch1)
	ch1 <- 1
	var input string
	fmt.Scanln(&input)
}

func print(s string, delay time.Duration, in chan int, out chan int) {
	for i := 0; i < 5; i++ {
		<-in
		println(s)
		time.Sleep(delay * time.Second)
		out <- 1
	}
}

/*
expected output:

Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/
