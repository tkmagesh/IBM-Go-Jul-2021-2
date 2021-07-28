package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(ch)
	for no := range ch {
		println(no)
	}
	fmt.Println("Done")
}

func fibonacci(ch chan int) {
	x, y := 0, 1
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
