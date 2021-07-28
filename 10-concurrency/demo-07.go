package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go add(10, 20, ch)
	result := <-ch
	fmt.Println(result)
	fmt.Println("Done")
}

func add(x, y int, ch chan int) {
	time.Sleep(5 * time.Second)

	//commenting the following lines will lead to a deadlock
	result := x + y
	ch <- result
}
